package handlers

import (
	"net/http"

	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: "1", Title: "1984", Author: "George Orwell"},
	{ID: "2", Title: "Brave New World", Author: "Aldous Huxley"},
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func EditBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func DeleteBook(c *gin.Context) {
	//find the book first
	//statusnotfound if not found
	//delete

	// id := c.Param("id")
	// c.JSON(http.StatusNotFound, gin.H{"message": "book not found", "ID": id})

	c.JSON(http.StatusNoContent, gin.H{"message": "book deleted"})
}
