package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Isbn string `json:"isbn"`
	AuthorID uint   // Foreign key
    Author Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
