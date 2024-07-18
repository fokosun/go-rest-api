package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Gravatar  string `json:"gravatar"`
	CreatedBy uint
	UpdatedBy uint
}

func (a *Author) SetCreatedBy(userId uint) error {
	a.CreatedBy = userId
	return nil
}

func (a *Author) SetUpdatedBy(userId uint) error {
	a.UpdatedBy = userId
	return nil
}

// Validate validates the User fields.
func (u *Author) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
