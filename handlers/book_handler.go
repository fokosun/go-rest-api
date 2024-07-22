package handlers

import (
	"net/http"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
	books := []models.Book{}
	config.DB.Find(&books)

	config.DB.Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Firstname", "Lastname", "Gravatar", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt")
	}).Find(&books)

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	var book models.Book
	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Book not found"})
		return
	}

	var qb models.Book

	config.DB.Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Firstname", "Lastname", "Gravatar", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt")
	}).First(&qb, book.ID)

	c.JSON(http.StatusOK, NewBook{ID: int(qb.ID), Title: qb.Title, Isbn: qb.Isbn, Author: qb.Author, CreatedAt: qb.CreatedAt, UpdatedAt: qb.UpdatedAt})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "title is required"})
		return
	}

	if book.UserID == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "user_id is required"})
		return
	}

	if book.AuthorID == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "author_id is required"})
		return
	}

	// also check if the author exist
	var bookAuthor models.Author
	if err := config.DB.First(&bookAuthor, book.AuthorID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Author not found"})
		return
	}

	config.DB.Create(&book)

	var qb models.Book

	config.DB.Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Firstname", "Lastname", "Gravatar", "CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt")
	}).First(&qb, book.ID)

	c.JSON(http.StatusCreated, NewBook{ID: int(qb.ID), Title: qb.Title, Isbn: qb.Isbn, Author: qb.Author, CreatedAt: qb.CreatedAt, UpdatedAt: qb.UpdatedAt})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	var user models.User

	if err := config.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Book not found"})
		return
	}

	if err := config.DB.Where("email = ?", c.MustGet("email")).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "There was a problem processing this request. Please try again."})
		return
	}

	if user.ID != book.UserID {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "You are not authorized to perform this action."})
		return
	}

	config.DB.Delete(&book)
	c.JSON(http.StatusOK, SuccessResponse{Message: "Book deleted"})
}
