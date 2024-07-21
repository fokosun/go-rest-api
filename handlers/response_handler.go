package handlers

import (
	"time"

	"github.com/fokosun/go-rest-api/models"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	ValidationErrorMessage string `json:"message"`
}

type NewUser struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewBook struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Isbn      string        `json:"isbn"`
	Author    models.Author `json:"author"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type LoginToken struct {
	Token string `json:"token"`
}
