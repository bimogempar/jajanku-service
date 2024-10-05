package user

import (
	"jajanku_service/dto"
	"jajanku_service/helpers"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Register(c *fiber.Ctx) error {
	req := new(dto.RegisterUserRequest)
	if err := c.BodyParser(req); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}

	user, err := h.service.Register(&User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	
	if err != nil {
		return helpers.APIError(c, fiber.StatusInternalServerError, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, user, "success register user")
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req := new(dto.LoginUserRequest)
	if err := c.BodyParser(req); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}
	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return helpers.APIError(c, fiber.StatusInternalServerError, err)
	}
	return helpers.APIResponse(c, fiber.StatusOK, token, "success login")
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		return helpers.APIError(c, fiber.StatusNotFound, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, user, "success get user")
}

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers()
	if err != nil {
		return helpers.APIError(c, fiber.StatusInternalServerError, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, users, "success get list user")
}
