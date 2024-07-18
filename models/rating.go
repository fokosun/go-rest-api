package models

type Rating struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	BookID  int
	Rating  int    `json:"rating" gorm:"default:1"`
	Comment string `json:"comment"`
}

func (r *Rating) SetBookID(bookID int) error {
	r.BookID = bookID
	return nil
}

func (r *Rating) SetUserID(userID uint) error {
	r.UserID = userID
	return nil
}
