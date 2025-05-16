package postgres

import (
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

func (s *PostgresStorage) DB() *gorm.DB {
	return s.db
}

func NewPostgresStorage(ctx context.Context, DBHost, DBUser, DBName, DBPass, DBPort string) (*PostgresStorage, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		DBHost,
		DBUser,
		DBName,
		DBPass,
		DBPort,
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
