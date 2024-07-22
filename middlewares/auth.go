package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("jwt-secret"))

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.Mode() == gin.TestMode {
			var testUser models.User

			testUser.Firstname = "test"
			testUser.Lastname = "last"
			testUser.Email = "test@example.com"
			testUser.Password = "validPass"
			testUser.SetPassword("validPass")

			if err := config.DB.Where("email = ?", testUser.Email).First(&testUser).Error; err != nil {
				// Save the user to the database
				config.DB.Create(&testUser)
			}

			// In test mode, bypass actual authentication
			c.Set("email", "test@example.com")
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Ensure the requesting user exists
		var user models.User
		userEmail := claims.Email
		if err := config.DB.Where("email = ?", userEmail).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired. Please login and try again"})
			c.Abort()
			return
		}

		// Token is valid, store user information in the context
		c.Set("email", userEmail)
		c.Next()
	}
}
