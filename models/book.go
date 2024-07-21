package models

import (
	"time"
)

type Book struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `json:"title"`
	Isbn      string `json:"isbn"`
	CreatedBy uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	AuthorID  uint   `gorm:"not null"` // Foreign key
	Author    Author `gorm:"-,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
