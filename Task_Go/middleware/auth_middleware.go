package middleware

import (
	"strings"
	"task_go/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.JSON(fiber.Map{"error": "No token"})
	}
	tokenString := strings.Split(token, " ")[1]
	jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})
	return c.Next()
}