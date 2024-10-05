package internal

import (
	"jajanku_service/internal/category"
	"jajanku_service/internal/config"
	"jajanku_service/internal/product"
	"jajanku_service/internal/user"

	"gorm.io/gorm"
)

type Registry struct {
	db   *gorm.DB
	conf *config.Config
}

func NewRegistry(db *gorm.DB, conf *config.Config) *Registry {
	return &Registry{
		db:   db,
		conf: conf,
	}
}

func (r *Registry) NewUserHandler() *user.Handler {
	repo := user.NewGormRepository(r.db)
	service := user.NewUserService(repo, r.conf.JWTSecret)
	return user.NewHandler(service)
}

func (r *Registry) NewProductHandler() *product.Handler {
	repo := product.NewGormRepository(r.db)
	service := product.NewProductService(repo)
	categoryService := r.InitCategory()
	return product.NewHandler(service, categoryService)
}

func (r *Registry) NewCategoryHandler() *category.Handler {
	service := r.InitCategory()
	return category.NewHandler(service)
}

func (r *Registry) InitCategory() category.Service {
	repo := category.NewGormRepository(r.db)
	return category.NewCategoryService(repo)
}

