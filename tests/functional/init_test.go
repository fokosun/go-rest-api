package tests

import (
	"os"
	"testing"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/fokosun/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var testUser models.User

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "pass")
	os.Setenv("DB_NAME", "books_store")
	os.Setenv("DB_PORT", "5432")

	config.ConnectDatabase()

	testUser.Firstname = "test"
	testUser.Lastname = "test"
	testUser.Email = "test@example.com"
	testUser.SetPassword("validpassword")

	// Save the user to the database
	config.DB.FirstOrCreate(&testUser)

	router = routes.SetupRouter()

	// Run tests
	code := m.Run()

	// Cleanup
	config.DB.Delete(&testUser)

	os.Exit(code)
}
