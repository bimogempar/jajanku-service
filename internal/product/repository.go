package product

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(Product *Product) error
	GetByID(id string) (*Product, error)
	GetByCategory(email string) ([]*Product, error)
	GetAll() ([]*Product, error)
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &GormRepository{db}
}

func (r *GormRepository) Create(Product *Product) error {
	Product.ID = uuid.New()
	return r.db.Create(Product).Error
}

func (r *GormRepository) GetByID(id string) (*Product, error) {
	var Product Product
	if err := r.db.First(&Product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &Product, nil
}

func (r *GormRepository) GetByCategory(email string) ([]*Product, error) {
	var Product []*Product
	result := r.db.Where("email = ?", email).First(&Product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Return nil, nil if the Product is not found
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return Product, nil
}

func (r *GormRepository) GetAll() ([]*Product, error) {
	var Products []*Product
	if err := r.db.Find(&Products).Error; err != nil {
		return nil, err
	}
	return Products, nil
}