package routes

import (
	"github.com/fokosun/go-rest-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Auth Routes
	auth := router.Group("/auth")
    {
        auth.POST("/login", handlers.Login)
        // auth.POST("/logout", handlers.Logout)
        // auth.POST("/invalidate-token", handlers.InvalidateToken)
    }

	// Users Routes
	users := router.Group("/users")
    {
        users.GET("/", handlers.GetUsers)
        users.GET("/:id", handlers.GetUserByID)
        users.POST("/", handlers.CreateUser)
        users.PUT("/:id", handlers.UpdateUser)
        users.DELETE("/:id", handlers.DeleteUser)
    }

	// Authors Routes
	// authors := router.Group("/authors")
    // {
    //     authors.GET("/", handlers.GetAuthors)
    //     authors.GET("/:id", handlers.GetAuthorByID)
    //     authors.POST("/", handlers.CreateAuthor)
    //     authors.PUT("/:id", handlers.EditAuthor)
    //     authors.DELETE("/:id", handlers.DeleteAuthor)
    // }

	// Books Routes
	books := router.Group("/books")
	{
		books.GET("/", handlers.GetBooks)
		books.GET("/:id", handlers.GetBookByID)
		books.POST("/", handlers.CreateBook)
		books.PUT("/:id", handlers.EditBook)
		books.DELETE("/:id", handlers.DeleteBook)

		//Ratings
		// books.GET("/ratings", handlers.GetRatings)
		// books.GET("/ratings/:id", handlers.GetRatingByBookID)
		// books.POST("/ratings", handlers.CreateRating)
		// books.PUT("/ratings/:id", handlers.EditRating)
		// books.DELETE("/ratings/:id", handlers.DeleteRating)
	}

	return router
}
