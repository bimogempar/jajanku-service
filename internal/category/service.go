package category

import (
	"errors"
)

type Service interface {
	CreateCategory(Category *Category) (*Category, error)
	GetCategory(id string) (*Category, error)
	ListCategorys() ([]*Category, error)
	GetCategoryByNames(name string) (*Category, error)
}

type CategoryService struct {
	repo      Repository
}

func NewCategoryService(repo Repository) Service {
	return &CategoryService{
		repo:      repo,
	}
}

func (s *CategoryService) CreateCategory(Category *Category) (*Category, error) {
	if _, err := s.GetCategoryByNames(Category.Name); err == nil {
		return nil, errors.New("Category already exists")
	}
	return Category, s.repo.Create(Category)
}

func (s *CategoryService) GetCategory(id string) (*Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) ListCategorys() ([]*Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetCategoryByNames(name string) (*Category, error) {
	return s.repo.GetByName(name)
}