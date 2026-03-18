package handlers

import (
	"net/http"
	"task_go/config"
	"task_go/models"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	var task models.Task
	if c.ShouldBindJSON(&task) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	config.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func GetTasksHandler(c *gin.Context) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if config.DB.First(&task, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	var input models.Task
	if c.ShouldBindJSON(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	task.Title = input.Title
	task.Completed = input.Completed
	config.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if config.DB.First(&task, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	config.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
