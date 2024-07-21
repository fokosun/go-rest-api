package models

import (
	"time"
)

type Book struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	CreatedAt time.Time
	UpdatedAt time.Time
	AuthorID  uint   `json:"author_id" gorm:"not null"` // Foreign key
	Author    Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
