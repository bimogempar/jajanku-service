package user

import "gorm.io/gorm"

type Repository interface {
	GetAllUser() ([]AllUsers, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) GetAllUser() ([]AllUsers, error) {
	var users []AllUsers
	// keknya bakalan fetching all user from firebase mul
	return users, nil
}
