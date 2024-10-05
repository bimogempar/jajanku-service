package product

import (
	"jajanku_service/dto"
	"jajanku_service/helpers"
	"jajanku_service/internal/category"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


type Handler struct {
	service Service
	categoryService category.Service
}

func NewHandler(service Service, categoryService category.Service) *Handler {
	return &Handler{service, categoryService}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req := new(dto.AddProductRequest)
	if err := c.BodyParser(req); err != nil {
		return helpers.APIError(c, fiber.StatusBadRequest, err)
	}

	cat, err := h.categoryService.GetCategoryByNames(req.Category)

	if cat == nil {
		cat, err = h.categoryService.CreateCategory(&category.Category{
			CategoryID: uuid.New(),
			Name: req.Category,
		})
	}


	Product, err := h.service.CreateProduct(&Product{
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		CategoryID: cat.CategoryID,
		Category: *cat,
	})

	productResponse := dto.ProductResponse{
		ID:       Product.ID,
		Name:     Product.Name,
		Price:    Product.Price,
		Stock:    Product.Stock,
		Category: Product.Category.Name,
	}
	
	if err != nil {
		return helpers.APIError(c, fiber.StatusInternalServerError, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, productResponse, "success register Product")
}


func (h *Handler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	Product, err := h.service.GetProduct(id)
	productResponse := dto.ProductResponse{
		ID:       Product.ID,
		Name:     Product.Name,
		Price:    Product.Price,
		Stock:    Product.Stock,
		Category: Product.Category.Name,
	}
	if err != nil {
		return helpers.APIError(c, fiber.StatusNotFound, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, productResponse, "success get Product")
}

func (h *Handler) ListProducts(c *fiber.Ctx) error {
	Products, err := h.service.ListProducts()
	if err != nil {
		return helpers.APIError(c, fiber.StatusInternalServerError, err)
	}

	return helpers.APIResponse(c, fiber.StatusOK, Products, "success get list Product")
}