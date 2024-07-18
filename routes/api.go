package routes

import (
	"github.com/fokosun/go-rest-api/handlers"
	"github.com/fokosun/go-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupApiRouter(router *gin.Engine) {
	// Auth Routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
	}

	// Register a new user
	router.POST("/register", handlers.CreateUser)

	// Users Routes
	users := router.Group("/users").Use(middlewares.AuthMiddleware())
	{
		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetUserByID)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)

		// A user i.e reader can create/view/update an author
		users.POST("/authors", handlers.CreateAuthor)
		users.GET("/authors", handlers.GetAuthors)
		users.GET("/authors/:id", handlers.GetAuthor)
		users.PUT("/authors/:id", handlers.EditAuthor)
	}

	// Books Routes
	books := router.Group("/books").Use(middlewares.AuthMiddleware())
	{
		books.GET("/", handlers.GetBooks)
		books.GET("/:id", handlers.GetBookByID)
		books.POST("/", handlers.CreateBook)
		books.PUT("/:id", handlers.EditBook)
		books.DELETE("/:id", handlers.DeleteBook)

		//Ratings
		// books.GET("/ratings", handlers.GetRatings)
		// books.GET("/ratings/:id", handlers.GetRatingByBookID)
		books.POST("/:id/ratings", handlers.CreateRating)
		// books.PUT("/ratings/:id", handlers.EditRating)
		// books.DELETE("/ratings/:id", handlers.DeleteRating)
	}
}
