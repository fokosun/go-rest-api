package handlers

import (
	"net/http"

	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

var rating models.Rating

func CreateRating(c *gin.Context) {
	// Bind the JSON input to the struct
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the Rating
    err := rating.Validate()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"validation_error": err.Error()})
		return
    }

	//ensure the given book id exists
	//set the user id from gin context
	//create the rating
}