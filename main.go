package main

import (
	"jajanku_service/internal/config"
	"jajanku_service/internal/modules/user"
	"jajanku_service/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf, err := config.New()
	app := fiber.New()

	db, err := config.InitDB(conf)
	if err != nil {
		log.Fatalf("failed connect db")
	}

	app.Get("/healthy", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]interface{}{
			"status":  200,
			"message": "Server running successfully",
		})
	})

	userHandler := user.InitUser(db, conf.JWTConfig.SecretKey)

	routes.UserRoutes(app, userHandler, conf.JWTConfig.SecretKey)

	log.Println("api running on localhost:3000")
	app.Listen(":3000")
}
