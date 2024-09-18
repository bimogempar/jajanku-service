package user

import (
	"jajanku_service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouteUser(app *fiber.App, handler *Handler) {
	userRouteGroup := app.Group("/api")
	userRouteGroup.Get("/users", middleware.JWTProtected(), handler.ListUsers)
	userRouteGroup.Post("/register", handler.Register)
	userRouteGroup.Post("/login", handler.Login)
}
