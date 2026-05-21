package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBUser     string
	DBName     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("load .env file: %w", err)
	}

	cfg := Config{
		Port:       os.Getenv("PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}

	if cfg.Port == "" {
		return Config{}, fmt.Errorf("missing PORT")
	}
	if cfg.DBUser == "" || cfg.DBName == "" || cfg.DBPassword == "" || cfg.DBHost == "" || cfg.DBPort == "" {
		return Config{}, fmt.Errorf("missing required database configuration")
	}

	return cfg, nil
}
