package main

import (
	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/fokosun/go-rest-api/routes"
)

// var ctx = context.Background()

func main() {
	Init()

	router := routes.SetupRouter()
	router.Run(":8080")
}

func Init() {
	config.ConnectDatabase()
	db := config.DB

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Rating{})

	// Add check constraint for Rating field
	// db.Exec("ALTER TABLE ratings ADD CONSTRAINT check_rating CHECK (rating IN (1, 2, 3, 4, 5))")

	// Set a key-value pair
	// err := config.Client.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	log.Fatalf("Error setting key: %v", err)
	// }
}
