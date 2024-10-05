package product

import (
	"jajanku_service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouteProduct(group fiber.Router, app *fiber.App, handler *Handler) {
	userRouteGroup := group.Group("/products", middleware.JWTProtected())
	userRouteGroup.Get("/", handler.ListProducts)
	userRouteGroup.Post("/create", handler.Create)
	userRouteGroup.Get("/:id", handler.GetProduct)
}