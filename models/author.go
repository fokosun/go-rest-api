package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Author struct {
	ID        uint   `gorm:"primarykey"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Gravatar  string `json:"gravatar"`
	CreatedBy uint
	UpdatedBy uint
	CreatedAt time.Time
	UpdatedAt time.Time
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
