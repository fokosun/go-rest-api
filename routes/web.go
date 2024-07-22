package routes

import (
	"github.com/fokosun/go-rest-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpWebRouter(router *gin.Engine) {
	// Auth Routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
	}

	// Register a new user
	router.POST("/register", handlers.RegisterUser)
}
