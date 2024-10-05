package category

import (
	"jajanku_service/helpers"

	"github.com/gofiber/fiber/v2"
	"jajanku_service/dto"
)


type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req := new(dto.AddCategoryRequest)
	if err := c.BodyParser(req); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}

	Category, err := h.service.CreateCategory(&Category{
		Name:     req.Name,
	})
	
	if err != nil {
		return helpers.APIError(c, fiber.StatusInternalServerError, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, Category, "success register Category")
}


func (h *Handler) GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	Category, err := h.service.GetCategory(id)
	CategoryResponse := dto.CategoryResponse{
		ID:       Category.CategoryID,
		Name:     Category.Name,
	}
	if err != nil {
		return helpers.APIError(c, fiber.StatusNotFound, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, CategoryResponse, "success get Category")
}

func (h *Handler) ListCategorys(c *fiber.Ctx) error {
	Categorys, err := h.service.ListCategorys()
	if err != nil {
		return helpers.APIError(c, fiber.StatusInternalServerError, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, Categorys, "success get list Category")
}