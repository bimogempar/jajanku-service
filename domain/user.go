package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Name     string    `json:name`
	Email    string    `json:email gorm:"unique"`
	Password string    `json:"-"`
	Role     string    `json:role`
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]*User, error)
}

type UserService interface {
	RegisterUser(user *User) error
	LoginUser(email, password string) (string, error)
	GetAllUsers() ([]*User, error)
}
