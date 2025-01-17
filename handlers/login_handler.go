package handlers

import (
	"net/http"

	"github.com/fokosun/go-rest-api/auth"
	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginDetails struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	reqPassword := loginDetails.Password

	var user models.User
	if err := config.DB.Where("email = ?", loginDetails.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid email or password"})
		return
	}

	if !user.CheckPassword(reqPassword) {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid email or password"})
		return
	}

	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginToken{Token: token})
}
