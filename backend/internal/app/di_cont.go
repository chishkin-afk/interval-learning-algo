package app

import (
	"context"
	"log/slog"
	"os"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/common/config"
	"github.com/chishkin-afk/interval-learning-algo/backend/internal/common/persistence/postgres"
	logger "github.com/chishkin-afk/interval-learning-algo/backend/pkg/log"
	"github.com/jmoiron/sqlx"
)

type DI struct {
	cfg *config.Config
	log *slog.Logger

	db *sqlx.DB
}

func (di *DI) Config() *config.Config {
	if di.cfg == nil {
		cfg, err := config.New(os.Getenv("APP_CONFIG_PATH"))
		if err != nil {
			slog.Error("can't load config",
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}

		di.cfg = cfg
	}

	return di.cfg
}

func (di *DI) Log() *slog.Logger {
	if di.log == nil {
		di.log = slog.New(logger.NewHandler(
			di.Config().App.Env))
	}

	return di.log
}

func (di *DI) DB() *sqlx.DB {
	if di.db == nil {
		db, err := postgres.Connect(&di.Config().Persistence)
		if err != nil {
			slog.Error("can't connect to db",
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}

		Add(func(ctx context.Context) error {
			errCh := make(chan error, 1)
			go func() {
				defer close(errCh)
				if err := db.Close(); err != nil {
					errCh <- err
				}
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case err, ok := <-errCh:
				if (ok) {
					return err
				}

				return nil
			}
		})

		di.db = db
	}

	return di.db
}