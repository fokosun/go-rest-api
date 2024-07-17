package main

import (
	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/fokosun/go-rest-api/routes"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Book{})

	router := routes.SetupRouter()
	router.Run(":8080")
}
