package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	UserID uint
	Books []Book `gorm:"foreignKey:AuthorID"`
}
