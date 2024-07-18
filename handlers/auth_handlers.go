package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Printf("Login")
}

func Logout(c *gin.Context) {
	fmt.Printf("Logout")
}

func InvalidateToken(c *gin.Context) {
	fmt.Printf("Invalidate Token")
}