package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"task_go/config"
	"task_go/models"
	"task_go/routes"
)

func main() {

	fmt.Println("Server running on port 4000")

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Task{})

	r := gin.Default()

	routes.RegisterRoutes(r)
	r.Run(":4000")
}
