package config

import (
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBName     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func LoadConfig() Config {
	cfg := Config{
		Port:       os.Getenv("PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
	return cfg
}
