package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_DSN     string
	JWT_SECRET string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	DB_DSN = os.Getenv("DB_DSN")
	JWT_SECRET = os.Getenv("JWT_SECRET")

	if DB_DSN == "" || JWT_SECRET == "" {
		log.Fatal("Missing environment variables")
	}
}