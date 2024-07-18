package models

import (
	"github.com/go-playground/validator/v10"
)

type Rating struct {
	ID      uint `gorm:"primaryKey"`
	UserID  int
	BookID  int `json:"book_id" validate:"required"`
	Rating  int `gorm:"default:1"`
	Comment string `json:"title"`
}

// Validate validates the Rating fields.
func (u *Rating) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
