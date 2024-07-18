package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Isbn string `json:"isbn"`
	AuthorID sql.NullInt64
}
