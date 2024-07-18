package models

type Rating struct {
	ID      uint `gorm:"primaryKey"`
	UserID  int
	BookID  int
	Rating  int    `gorm:"default:1"`
	Comment string `json:"title"`
}
