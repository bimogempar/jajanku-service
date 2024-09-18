package main

import (
	"jajanku_service/internal"
	"jajanku_service/internal/config"
	"jajanku_service/internal/user"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.LoadConfig()
	app := fiber.New()

	db, err := config.InitDB(conf)
	if err != nil {
		log.Fatalf("failed connect db", err)
	}

	registry := internal.NewRegistry(db, conf)
	if err != nil {
		log.Fatalf("failed load registry", err)
	}

	userHandler := registry.NewUserHandler()
	user.RegisterRouteUser(app, userHandler)

	log.Println("api running on localhost:3000")
	app.Get("/api/healthy", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]interface{}{
			"status":  200,
			"message": "Server running successfully",
		})
	})
	app.Listen(":3000")
}
