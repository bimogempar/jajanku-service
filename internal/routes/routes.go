package routes

import (
	"jajanku_service/internal"
	"jajanku_service/internal/product"
	"jajanku_service/internal/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, registry *internal.Registry) {
	
	apiGroup := app.Group("/api")
	userHandler := registry.NewUserHandler()
	productHandler := registry.NewProductHandler()

	user.RegisterRouteUser(apiGroup, app, userHandler)
	product.RegisterRouteProduct(apiGroup, app, productHandler)
}
