package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Log      Log
		Database Database
	}

	App struct {
		Name string `env:"APP_NAME,required"`
		Port string `env:"APP_PORT,required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	Database struct {
		Host     string `env:"DB_HOST,required"`
		Port     string `env:"DB_PORT,required"`
		User     string `env:"DB_USER,required"`
		Password string `env:"DB_PASSWORD,required"`
		Name     string `env:"DB_NAME,required"`
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
		log.Println("Warning: .env file not found, using environment variables")
	}

	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
