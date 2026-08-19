package postgres

import (
	"fmt"
	"path/filepath"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/common/config"
	"github.com/golang-migrate/migrate/v4"
)

func Migrate(cfg *config.Persistence) error {
	migrationsPath := filepath.Clean(cfg.MigrationsPath)
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
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.Postgres.Auth.User,
		cfg.Postgres.Auth.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Auth.DB,
	)
}