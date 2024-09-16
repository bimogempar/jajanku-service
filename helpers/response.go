package helpers

import (
	"github.com/gofiber/fiber/v2"
)

func APIResponse(c *fiber.Ctx, statusCode int, data interface{}, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
}

func APIError(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode,
		"message": "error",
		"error":   err.Error(),
	})
}
