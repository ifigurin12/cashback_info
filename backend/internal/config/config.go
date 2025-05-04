package config

import (
	"fmt"
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
	Env        `env:"ENV" env-required:"true" validate:"required,oneof=local dev prod"`
	Storage    `env-required:"true"`
	HTTPServer `env-required:"true"`
}

type Storage struct {
	Host string `env:"DB_HOST" env-required:"true"`
	Port string `env:"DB_PORT" env-required:"true"`
	User string `env:"DB_USER" env-required:"true"`
	Pass string `env:"DB_PASS" env-required:"true"`
	Name string `env:"DB_NAME" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `env:"HTTP_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
}

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("error reading environment variables: %v", err)
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	if cfg.Env == "" {
		return fmt.Errorf("env is required")
	}
	if cfg.Storage.Host == "" || cfg.Storage.Port == "" || cfg.Storage.User == "" || cfg.Storage.Pass == "" || cfg.Storage.Name == "" {
		return fmt.Errorf("all storage fields are required")
	}
	if cfg.HTTPServer.Address == "" {
		cfg.HTTPServer.Address = "localhost:8080"
	}
	if cfg.HTTPServer.Timeout == 0 {
		cfg.HTTPServer.Timeout = 5 * time.Second
	}
	if cfg.HTTPServer.IdleTimeout == 0 {
		cfg.HTTPServer.IdleTimeout = 60 * time.Second
	}

	if cfg.Env != EnvLocal && cfg.Env != EnvDev && cfg.Env != EnvProd {
		return fmt.Errorf("invalid env value: %s", cfg.Env)
	}

	return nil
}
