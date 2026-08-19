package postgres

import (
	"fmt"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/common/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Persistence) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", getDataSourceName(cfg))
	if err != nil {
		return nil, fmt.Errorf("can't open conn with pg: %w", err)
	}

	applySettings(db, cfg)
	return db, nil	
}

func applySettings(db *sqlx.DB, cfg *config.Persistence) {
	db.SetConnMaxIdleTime(cfg.Postgres.Conns.MaxIdleTime)
	db.SetConnMaxLifetime(cfg.Postgres.Conns.MaxLifetime)
	db.SetMaxIdleConns(cfg.Postgres.Conns.MaxIdles)
	db.SetMaxOpenConns(cfg.Postgres.Conns.MaxOpens)
}

func getDataSourceName(cfg *config.Persistence) string {
	return fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode,
		cfg.Postgres.Auth.User,
		cfg.Postgres.Auth.Password,
		cfg.Postgres.Auth.DB,
	)
}