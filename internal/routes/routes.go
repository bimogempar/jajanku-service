package routes

import (
	"jajanku_service/internal"
	"jajanku_service/internal/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, registry *internal.Registry) {
	userHandler := registry.NewUserHandler()
	user.RegisterRouteUser(app, userHandler)
}
