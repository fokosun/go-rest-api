package handlers

import (
	"net/http"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

var author models.Author
var user models.User

func CreateAuthor(c *gin.Context) {
	// Bind the JSON input to the struct
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Validate the Author
    err := author.Validate()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"validation_error": err.Error()})
		return
    }

	userEmail := c.MustGet("email").(string)
	if err := config.DB.Where("email = ?", userEmail).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "There was a problem processing this request. Please try again."})
		return
	}

	if err = author.SetCreatedBy(user.ID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	// Save the author to the database
    result := config.DB.Create(&author)
    if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error creating author:":result.Error})
        return
    }

	c.JSON(http.StatusCreated, author)
}

func GetAuthors(c *gin.Context) {
	authors := []models.Author{}
	config.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func GetAuthor(c *gin.Context) {
	var author models.Author
	if err := config.DB.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

func EditAuthor(c *gin.Context) {
	var author models.Author
	if err := config.DB.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userEmail := c.MustGet("email").(string)
	if err := config.DB.Where("email = ?", userEmail).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "There was a problem processing this request. Please try again."})
		return
	}

	if err := author.SetUpdatedBy(user.ID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	config.DB.Save(&author)
	c.JSON(http.StatusOK, author)
}