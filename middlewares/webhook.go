package middlewares

import (
	"github.com/gin-gonic/gin"
)

func WebhookMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
