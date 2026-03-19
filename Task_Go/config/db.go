package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() {
	db, err := gorm.Open(postgres.Open(DB_DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	DB = db
	log.Println("Database connected")
}
