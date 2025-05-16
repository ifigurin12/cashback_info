package main

import (
	"cashback_info/internal/config"
	familyuser "cashback_info/internal/model/family/user"
	familyuserrepo "cashback_info/internal/repository/family/user"
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

	storage, err := postgres.NewPostgresStorage(context.Background(), cfg.Host, cfg.User, cfg.Name, cfg.Pass, cfg.Port)
	if err != nil {
		logger.Error("Failed to connect to Postgres:", err)
		os.Exit(1)
	}

	familyUserRepository := familyuserrepo.NewFamilyUserRepository(storage.DB())
	err = familyUserRepository.Create(familyuser.FamilyUser{FamilyID: "365afedc-e49d-4ab7-a243-158f360dd11c", UserID: "59e807df-1d32-4c3f-9890-5aa716de87c6"})
	if err != nil {
		logger.Error("Failed to create family user:", err)
		os.Exit(1)
	}

	logger.Info("Connected to Postgres")
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
