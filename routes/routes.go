package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	env := os.Getenv("ENV")

	if env == "development" {
		// Trust all proxies (not recommended for production)
		err := router.SetTrustedProxies(nil)
		if err != nil {
			panic(err)
		}
	}

	SetupApiRouter(router)
	SetUpWebRouter(router)
	SetUpWebhookRouter(router)

	return router
}
