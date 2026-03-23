package handlers

import (
	"fmt"
	"task_go/config"
	"task_go/models"

	"github.com/gofiber/fiber/v2"
)

var taskQueue = make(chan models.Task, 100)

func taskWorker() {
	for task := range taskQueue {
		fmt.Println("Background Processing Task:", task.ID, task.Title)

	}
}

func init() {
	go taskWorker()
}

func CreateTaskHandler(c *fiber.Ctx) error {
	var task models.Task
	c.BodyParser(&task)

	config.DB.Create(&task)

	go func(t models.Task) {
		taskQueue <- t
	}(task)

	return c.JSON(task)
}

func GetTasksHandler(c *fiber.Ctx) error {
	var tasks []models.Task
	config.DB.Find(&tasks)
	return c.JSON(tasks)
}

func UpdateTaskHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	config.DB.First(&task, id)

	var input models.Task
	c.BodyParser(&input)

	task.Title = input.Title
	task.Completed = input.Completed
	config.DB.Save(&task)

	go func(t models.Task) {
		taskQueue <- t
	}(task)

	return c.JSON(task)
}

func DeleteTaskHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	config.DB.First(&task, id)

	config.DB.Delete(&task)

	go func(t models.Task) {
		taskQueue <- t
	}(task)

	return c.JSON(fiber.Map{
		"message": "Deleted",
	})
}
