package handlers

import (
	"task_go/config"
	"task_go/models"
	"github.com/gofiber/fiber/v2"
)

func CreateTaskHandler(c *fiber.Ctx) error {
	var task models.Task
	c.BodyParser(&task)
	config.DB.Create(&task)
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

	return c.JSON(task)
}

func DeleteTaskHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	config.DB.First(&task, id)
	config.DB.Delete(&task)
	return c.JSON(fiber.Map{
		"message": "Deleted",
	})
}