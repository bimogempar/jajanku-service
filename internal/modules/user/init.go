package user

import "gorm.io/gorm"

func InitUser(db *gorm.DB, jwtSecret string) *UserHandler {
	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo, jwtSecret)
	return NewUserHandler(userService)
}
