package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/maulanarisqimustofa/jajanku-project/app/config"
	"github.com/maulanarisqimustofa/jajanku-project/app/middlewares"
	"github.com/maulanarisqimustofa/jajanku-project/modules/entities/user"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB) *fiber.App {
	fb := config.SetupFirebase()
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "server running successfully",
		})
	})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewUserHandler(userService)

	v1 := app.Group("/api/v1", middlewares.FirebaseAuth(fb))
	v1.Get("/users", userHandler.GetAllUsers)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "404 not found",
		})
	})

	return app
}
