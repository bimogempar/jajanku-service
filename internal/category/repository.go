package category

import "gorm.io/gorm"

type Repository interface {
	Create(Product *Category) error
	GetByID(id string) (*Category, error)
	GetAll() ([]*Category, error)
	GetByName(name string) (*Category, error)
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &GormRepository{db}
}

func (r *GormRepository) Create(Category *Category) error {
	return r.db.Create(Category).Error
}

func (r *GormRepository) GetByID(id string) (*Category, error) {
	var Category Category
	if err := r.db.First(&Category, "category_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &Category, nil
}

func (r *GormRepository) GetAll() ([]*Category, error) {
	var Categories []*Category
	if err := r.db.Find(&Categories).Error; err != nil {
		return nil, err
	}
	return Categories, nil
}

func (r *GormRepository) GetByName(name string) (*Category, error) {
	var Category Category
	if err := r.db.First(&Category, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &Category, nil
}