package user

import (
	"jajanku_service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouteUser(group fiber.Router, app *fiber.App, handler *Handler) {
	userRouteGroup := group.Group("/users")
	userRouteGroup.Get("/", middleware.JWTProtected(), handler.ListUsers)
	userRouteGroup.Post("/register", handler.Register)
	userRouteGroup.Post("/login", handler.Login)
	userRouteGroup.Get("/:id", handler.GetUser)
}
