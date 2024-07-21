package handlers

import "time"

type ErrorResponse struct {
	Message string `json:"message"`
}

type NewUser struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginToken struct {
	Token string `json:"token"`
}
