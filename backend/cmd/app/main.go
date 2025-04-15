package main

import (
	"cashback_info/internal/config"
	"cashback_info/internal/storage/postgres"
	"context"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	logger := setupLogger(cfg)
	logger.Info("cashback-info service started up in %s mode", cfg.Env)

	storage, err := postgres.NewPostgresStorage(context.Background(), cfg.Storage)
	if err != nil {
		logger.Error("Failed to connect to Postgres:", err)
		os.Exit(1)
	}

	defer storage.Close()

}

func setupLogger(cfg *config.Config) *slog.Logger {
	var logger *slog.Logger

	switch cfg.Env {
	case config.EnvLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvDev:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}
