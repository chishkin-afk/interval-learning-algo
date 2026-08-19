package main

import (
	"log/slog"

	"github.com/chishkin-afk/interval-learning-algo/backend/internal/app"
	"github.com/chishkin-afk/interval-learning-algo/backend/internal/common/persistence/postgres"
)

func main() {
	var di app.DI
	di.Log().Info("start migrations...", 
		slog.String("migrations_path", di.Config().Persistence.MigrationsPath))
	
	if err := postgres.Migrate(&di.Config().Persistence); err != nil {
		di.Log().Error("can't migrate db",
			slog.String("error", err.Error()),
		)
	}

	di.Log().Info("all migrations has been applied.")
}