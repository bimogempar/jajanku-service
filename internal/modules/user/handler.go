package user

import (
	"jajanku_service/domain"
	"jajanku_service/dto"
	"jajanku_service/helpers"
	"jajanku_service/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var req dto.RegisterUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}
	if err := pkg.Validate.Struct(&req); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}
	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.UserService.RegisterUser(user); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, user, "success register")
}

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	var req dto.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}
	if err := pkg.Validate.Struct(&req); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}
	token, err := h.UserService.LoginUser(req.Email, req.Password)
	if err != nil {
		return helpers.APIError(c, fiber.StatusUnauthorized, err)
	}
	return helpers.APIResponse(c, fiber.StatusOK, token, "success")
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		return helpers.APIError(c, fiber.StatusNotFound, err)
	}
	return helpers.APIResponse(c, fiber.StatusOK, users, "success")
}
