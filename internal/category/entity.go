package category

import "github.com/google/uuid"

type Category struct {
	CategoryID uuid.UUID `json:"category_id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name       string    `json:"name"`
}