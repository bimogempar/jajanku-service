package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Name     string    `json:"name"` 
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
	Role     string    `json:"role"`
}
