package config

import (
	"errors"
	"log"
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
	PostgresDSN string `yaml:"postgres_dsn" env-default:"postgres://postgres:postgres@localhost:5432/task_tracker?sslmode=disable"`
}

type Config struct {
	Env        Environment      `yaml:"env" env-default:"local"`
	HTTPServer HTTPServerConfig `yaml:"http_server"`
	Database   DBConfig         `yaml:"db"`
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		log.Printf("Config path is empty")
		return nil, errors.New("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Printf("Config file not found: %s", configPath)
		return nil, err
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Printf("Failed to load config: %v", err)
		return nil, err
	}

	return &config, nil
}
