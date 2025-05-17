package main

import (
	"cashback_info/internal/config"
	authHandler "cashback_info/internal/handler/auth"
	bankHandler "cashback_info/internal/handler/bank"
	categoryHandler "cashback_info/internal/handler/category"
	userHandler "cashback_info/internal/handler/user"
	bankrepo "cashback_info/internal/repository/bank"
	categoryrepo "cashback_info/internal/repository/category"
	userrepo "cashback_info/internal/repository/user"
	"cashback_info/internal/router"
	authservice "cashback_info/internal/service/auth"
	bankservice "cashback_info/internal/service/bank"
	categoryservice "cashback_info/internal/service/category"
	"cashback_info/internal/service/password"
	"cashback_info/internal/service/token"
	userservice "cashback_info/internal/service/user"
	"cashback_info/internal/storage/postgres"
	"context"
	"log"
	"log/slog"
	"os"
)

// @title Cashback-info API
// @version 1.0
// @description Cashback-info API
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

	userRepository := userrepo.NewUserRepository(storage.DB())
	categoryRepository := categoryrepo.NewCategoryRepository(storage.DB())
	bankRepository := bankrepo.NewBankRepository(storage.DB())

	tokenService := token.NewTokenServiceImpl(cfg.JWT.SecretKey)
	passwordService := password.NewPasswordService()

	userService := userservice.NewUserService(userRepository, passwordService)
	categoryservice := categoryservice.NewCategoryService(categoryRepository)
	bankService := bankservice.NewBankService(bankRepository)
	authService := authservice.NewAuthService(userRepository, tokenService, passwordService)

	r := router.SetupRouter()

	userHandler.NewUserHandler(userService).SetupRoutes(r)
	authHandler.NewAuthHandler(authService).SetupRoutes(r)
	categoryHandler.NewCategoryHandler(categoryservice).SetupRoutes(r)
	bankHandler.NewBankHandler(bankService).SetupRoutes(r)

	r.Run()

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
