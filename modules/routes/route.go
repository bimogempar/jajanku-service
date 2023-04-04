package routes

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func InitRoute() {
	godotenv.Load(".env")

	app := fiber.New()
	app.Use(logger.New())

	v1 := app.Group("/api/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "404 not found",
		})
	})

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
