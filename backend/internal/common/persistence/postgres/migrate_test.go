package postgres

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/common/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type Auth struct {
	User string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	DB string `yaml:"db" validate:"required"`
}

func TestGetDatabaseUrl(t *testing.T) {
	tests := []struct {
		name     string
		cfg      *config.Persistence
		expected string
	}{
		{
			name: "standard config",
			cfg: &config.Persistence{
				Postgres: config.Postgres{
					Host:    "localhost",
					Port:    5432,
					SSLMode: "disable",
					Auth: Auth{
						User:     "testuser",
						Password: "testpass",
						DB:       "testdb",
					},
				},
			},
			expected: "postgresql://testuser:testpass@localhost:5432/testdb?sslmode=disable",
		},
		{
			name: "different port and ssl",
			cfg: &config.Persistence{
				Postgres: config.Postgres{
					Host:    "db.example.com",
					Port:    5433,
					SSLMode: "require",
					Auth: Auth{
						User:     "admin",
						Password: "s3cret",
						DB:       "production_db",
					},
				},
			},
			expected: "postgresql://admin:s3cret@db.example.com:5433/production_db?sslmode=require",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getDatabaseUrl(tc.cfg)
			if result != tc.expected {
				t.Errorf("getDatabaseUrl() = %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestMigrateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"sha256:22c89fe0d0f507606260237fd55e51f6137f58b2d5bcf6152242b96d9fe8f9a4",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	require.NoError(t, err, "failed to start postgres container")
	defer pgContainer.Terminate(ctx)

	host, err := pgContainer.Host(ctx)
	require.NoError(t, err)
	port, err := pgContainer.MappedPort(ctx, "5432/tcp")
	require.NoError(t, err)

	migrationsDir, err := filepath.Abs("../../../../migrations")
	require.NoError(t, err)

	cfg := &config.Persistence{
		MigrationsPath: migrationsDir,
		Postgres: config.Postgres{
			Host:    host,
			Port:    int(port.Num()),
			SSLMode: "disable",
			Auth: Auth{
				User:     "testuser",
				Password: "testpass",
				DB:       "testdb",
			},
		},
	}

	err = Migrate(cfg)

	assert.NoError(t, err, "Migrate should not return an error")
}