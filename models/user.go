package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type User struct {
	gorm.Model
	Firstname    string `json:"firstname" validate:"required"`
	Lastname     string `json:"lastname" validate:"required"`
	Email        string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

// CheckPassword compares the given password with the stored password hash.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

// Validate validates the User fields.
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
