package handlers

import (
	"task_go/config"
	"task_go/models"
	"task_go/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)

	config.DB.Create(&user)
	return c.JSON(fiber.Map{
		"message": "User registered",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var input models.User
	var user models.User
	c.BodyParser(&input)
	config.DB.Where("username = ?", input.Username).First(&user)

	if user.Password != input.Password {
		return c.JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}
	token, _ := utils.GenerateToken(user.ID)
	return c.JSON(fiber.Map{
		"token": token,
	})
}