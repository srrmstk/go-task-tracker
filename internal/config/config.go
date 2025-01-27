package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Environment string

const (
	Local Environment = "local"
	Dev   Environment = "dev"
	Prod  Environment = "prod"
)

type HTTPServerConfig struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DBConfig struct {
	PostgresDSN string `yaml:"postgres_dsn"`
}

type Config struct {
	Env        Environment      `yaml:"env" env-default:"local"`
	HTTPServer HTTPServerConfig `yaml:"http_server"`
	Database   DBConfig         `yaml:"database"`
}

const (
	ErrConfigPathEmpty = "config path is empty"
	ErrConfigNotFound  = "config file not found"
	ErrFailedToLoad    = "failed to load config"
)

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, errors.New(ErrConfigPathEmpty)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%v: %s", ErrConfigNotFound, configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, fmt.Errorf("%v: %v", ErrFailedToLoad, err)
	}

	return &config, nil
}
