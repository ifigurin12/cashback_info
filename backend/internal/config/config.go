package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal Env = "local"
	EnvDev   Env = "dev"
	EnvProd  Env = "prod"
)

type Env string

type Config struct {
	Env        `yaml:"env" env:"ENV" env-required:"true" validate:"required,oneof=local dev prod"`
	Storage    `yaml:"storage" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
}

type Storage struct {
	Host string `yaml:"host" env:"DB_HOST" env-required:"true"`
	Port string `yaml:"port" env:"DB_PORT" env-required:"true"`
	User string `yaml:"user" env:"DB_USER" env-required:"true"`
	Pass string `yaml:"pass" env:"DB_PASS" env-required:"true"`
	Name string `yaml:"name" env:"DB_NAME" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"HTTP_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
}

func New() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		return nil, fmt.Errorf("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
