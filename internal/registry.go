package internal

import (
	"jajanku_service/internal/config"
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
