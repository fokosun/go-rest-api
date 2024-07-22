package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetRatings(c *gin.Context) {
	ratings := []models.Rating{}
	config.DB.Find(&ratings)
	c.JSON(http.StatusOK, ratings)
}

func GetRatingsByBookID(c *gin.Context) {
	ratings := []models.Rating{}
	if err := config.DB.Where("book_id = ?", c.Param("id")).Find(&ratings).Error; err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Book does not exist"})
		return
	}

	c.JSON(http.StatusOK, ratings)
}

func CreateOrUpdateRating(c *gin.Context) {
	var rating models.Rating
	var book models.Book
	var user models.User

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "An unknown error occured. please try again."})
		return
	}

	// If Rating dont exists create new
	if err := config.DB.Where(&models.Rating{BookID: bookID, UserID: user.ID}).First(&rating).Error; err != nil {
		fmt.Println("Creating new Rating")

		// Bind the JSON input to the struct
		if err := c.ShouldBindJSON(&rating); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}

		// ensure the given book id exists
		if err := config.DB.First(&book, bookID).Error; err != nil {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: "Book not found"})
			return
		}

		// set the book id from gin context
		if err := rating.SetBookID(bookID); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
			return
		}

		// set the user id from gin context
		if err := rating.SetUserID(user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
			return
		}

		config.DB.Create(&rating)

		c.JSON(http.StatusCreated, rating)

		return
	}

	// Bind the JSON input to the struct
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	fmt.Println("Updating existing Rating")

	config.DB.Save(&rating)

	c.JSON(http.StatusOK, rating)
}
