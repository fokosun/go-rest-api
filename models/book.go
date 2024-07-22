package models

import (
	"time"
)

type Book struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	UserID    uint   `gorm:"not null"` // Foreign key
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	AuthorID  uint   `gorm:"not null"` // Foreign key
	Author    Author `gorm:"-,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
