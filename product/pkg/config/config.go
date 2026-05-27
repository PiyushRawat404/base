package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port       string
	DBName     string
	DBUser     string
	DBPort     string
	DBPassword string
	DBHost     string
	JWTSecret  string
}

func LoadConfig() (Env, error) {
	_ = godotenv.Load()

	cfg := Env{
		Port:       getenv("PORT", "8080"),
		DBName:     getenvAny("DB_NAME", "DBNAME"),
		DBUser:     getenvAny("DB_USER", "DBUSER"),
		DBPort:     getenvAnyDefault("5432", "DB_PORT", "DBPORT"),
		DBPassword: getenvAny("DB_PASSWORD", "DBPASSWORD"),
		DBHost:     getenvAnyDefault("localhost", "DB_HOST", "DBHOST"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}

	if cfg.DBName == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.JWTSecret == "" {
		return Env{}, errors.New("missing required environment variables")
	}

	return cfg, nil
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getenvAny(keys ...string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return ""
}

func getenvAnyDefault(fallback string, keys ...string) string {
	if value := getenvAny(keys...); value != "" {
		return value
	}
	return fallback
}
