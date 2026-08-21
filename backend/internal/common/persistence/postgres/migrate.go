package postgres

import (
	"fmt"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/common/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(cfg *config.Persistence) error {
	migrationsPath := "file://" + cfg.MigrationsPath
	migrate, err := migrate.New(migrationsPath, getDatabaseUrl(cfg))
	if err != nil {
		return fmt.Errorf("can't create migration: %w", err)
	}

	if err := migrate.Up(); err != nil {
		return fmt.Errorf("can't up migration: %w", err)
	}

	return nil
}

func getDatabaseUrl(cfg *config.Persistence) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Postgres.Auth.User,
		cfg.Postgres.Auth.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Auth.DB,
		cfg.Postgres.SSLMode,
	)
}