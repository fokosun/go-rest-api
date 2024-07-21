package models

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	AuthorID  sql.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
}
