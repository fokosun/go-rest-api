package models

import "time"

type Rating struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint
	BookID    int
	Rating    int    `json:"rating" gorm:"default:1"`
	Comment   string `json:"comment"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *Rating) SetBookID(bookID int) error {
	r.BookID = bookID
	return nil
}

func (r *Rating) SetUserID(userID uint) error {
	r.UserID = userID
	return nil
}
