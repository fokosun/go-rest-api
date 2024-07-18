package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func CreateRating(c *gin.Context) {
	fmt.Printf("Params %v\n", c.Query("comment"))
	var rating models.Rating
	var book models.Book
	var user models.User

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Ensure the requesting user exists
	userEmail := c.MustGet("email").(string)
	if err := config.DB.Where("email = ?", userEmail).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "There was a problem processing this request. Please try again."})
		return
	}

	// If Rating dont exists create new
	if err := config.DB.Where(&models.Rating{BookID:bookID, UserID: user.ID}).First(&rating).Error; err != nil {
		fmt.Printf("Creating new Rating")

		// Bind the JSON input to the struct
		if err := c.ShouldBindJSON(&rating); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// ensure the given book id exists
		if err := config.DB.First(&book, bookID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		//set the book id from gin context
		if err := rating.SetBookID(bookID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//set the user id from gin context
		if err := rating.SetUserID(user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		config.DB.Create(&rating)

		c.JSON(http.StatusOK, rating)

		return
	}

	fmt.Printf("Updating existing Rating %v\n", rating)

	// Bind the JSON input to the struct
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&rating)

	c.JSON(http.StatusOK, rating)
}
