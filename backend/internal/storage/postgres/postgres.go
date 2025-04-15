package postgres

import (
	"cashback_info/internal/config"
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	db *gorm.DB
}

func (s *PostgresStorage) Close() {
	postgresDB, _ := s.db.DB()
	postgresDB.Close()
}

func NewPostgresStorage(ctx context.Context, storageConfig config.Storage) (*PostgresStorage, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		storageConfig.Host,
		storageConfig.User,
		storageConfig.Name,
		storageConfig.Pass,
		storageConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB from GORM: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &PostgresStorage{db: db}, nil
}
