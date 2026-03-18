package routes

import (
	"task_go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/tasks", handlers.GetTasksHandler)
	r.POST("/tasks", handlers.CreateTaskHandler)
	r.PUT("/tasks/:id", handlers.UpdateTaskHandler)
	r.DELETE("/tasks/:id", handlers.DeleteTaskHandler)

}
