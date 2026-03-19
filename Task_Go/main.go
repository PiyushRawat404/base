package main

import (
	"task_go/config"
	"task_go/models"
	"task_go/routes"
	"task_go/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadConfig()
	utils.InitLogger()
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Task{}, &models.User{})

	app := fiber.New()

	routes.RegisterRoutes(app)

	utils.Logger.Info().Msg("Server running on port 4000")

	app.Listen(":4000")
}