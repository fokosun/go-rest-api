package handlers

import (
	"net/http"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found."})
		return
	}
	c.JSON(http.StatusOK, user)
}

func RegisterUser(c *gin.Context) {
	var user models.User

	// Bind the JSON input to the struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// Validate the user
	err := user.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// Validate the password length
	minLength := models.MinPasswordLength
	invalidPasswordLengthMessage := models.InvalidPasswordLengthMessage

	if !user.ValidatePassword(user.Password, minLength) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: invalidPasswordLengthMessage})
		return
	}

	if err = user.SetPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	// Save the user to the database
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "User already exists."})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found."})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if len(user.Password) > 0 {
		if err := user.SetPassword(user.Password); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}
	}

	// Update the user in the database
	if err := config.DB.Model(&user).Updates(user).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, user)
}

// ensure cascade on delete actually happens
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found."})
		return
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusNoContent, nil)
}
