package dto

import "github.com/google/uuid"

type AddProductRequest struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
	Category string `json:"category" validate:"required"`
}

type UpdateProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type ProductResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	Category string `json:"category"`
}