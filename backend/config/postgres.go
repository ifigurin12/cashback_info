package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context) (*pgxpool.Pool, error) {
	host := os.Getenv("POSTGRES_DB_HOST")
	user := os.Getenv("POSTGRES_DB_USER")
	password := os.Getenv("POSTGRES_DB_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB_NAME")
	port := os.Getenv("POSTGRES_DB_PORT")

	// Формирование строки подключения
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=verify-full TimeZone=Europe/Moscow", host, user, password, dbname, port)

	// Парсинг конфигурации подключения
	dbconfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	// Подключение к базе данных
	dbpool, err := pgxpool.NewWithConfig(ctx, dbconfig)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
