package routes

import (
	"task_go/handlers"
	"task_go/middleware"
	"github.com/gofiber/fiber/v2"
)
func RegisterRoutes(app *fiber.App) {
	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.LoginUser)
	app.Get("/tasks", middleware.AuthMiddleware, handlers.GetTasksHandler)
	app.Post("/tasks", middleware.AuthMiddleware, handlers.CreateTaskHandler)
	app.Put("/tasks/:id", middleware.AuthMiddleware, handlers.UpdateTaskHandler)
	app.Delete("/tasks/:id", middleware.AuthMiddleware, handlers.DeleteTaskHandler)
}
