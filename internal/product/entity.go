package product

import (
	"github.com/google/uuid"
	"jajanku_service/internal/category"
)

type Product struct {
	ID    uuid.UUID    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	Category 	category.Category `json:"category"`
	CategoryID 	uuid.UUID `gorm:"foreignKey:CategoryID"`
}