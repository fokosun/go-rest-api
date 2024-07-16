package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/fokosun/go-rest-api/handlers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/books", handlers.GetBooks)
    router.GET("/books/:id", handlers.GetBookByID)
    router.POST("/books", handlers.CreateBook)

    return router
}
