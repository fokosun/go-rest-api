package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type User struct {
	gorm.Model
	Firstname    string `json:"firstname" validate:"required"`
	Lastname     string `json:"lastname" validate:"required"`
	Email        string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password     string `json:"password,omitempty" validate:"required" gorm:"-"`
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

// ValidatePassword checks if the password meets the minimum length requirement
func (u *User) ValidatePassword(password string, minLength int) bool {
	return len(password) >= minLength
}

// Validate validates the User fields.
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// BeforeUpdate is a GORM hook that prevents the email field from being updated
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	var oldUser User
	if err := tx.First(&oldUser, u.ID).Error; err != nil {
		return err
	}

	if u.Email != oldUser.Email {
		return errors.New("email field cannot be updated")
	}

	return nil
}
