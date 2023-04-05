package user

import "github.com/gofiber/fiber/v2"

type userHandler struct {
	userService Service
}

func NewUserHandler(userService Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAllUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "GetAllUsers",
	})
}
