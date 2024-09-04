package routes

import (
	"jajanku_service/internal/middleware"
	"jajanku_service/internal/modules/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *user.UserHandler, JwtSecretKey string) {
	app.Post("/register", handler.RegisterUser)
	app.Post("/login", handler.LoginUser)
	app.Get("/users", middleware.JWTProtected(JwtSecretKey), handler.GetAllUsers)
}
