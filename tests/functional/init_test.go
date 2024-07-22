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
var testAuthor models.Author
var testBook models.Book

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "pass")
	os.Setenv("DB_NAME", "books_store")
	os.Setenv("DB_PORT", "5432")

	config.ConnectDatabase()

	testUser.Firstname = "Test User Firstname"
	testUser.Lastname = "Test User Lastname"
	testUser.Email = "test@example.com"
	testUser.SetPassword("validpassword")

	// Save the user to the database
	config.DB.FirstOrCreate(&testUser)

	testAuthor.Firstname = "Test Author Firstname"
	testAuthor.Lastname = "Test Author lastname"
	testAuthor.CreatedBy = testUser.ID

	// Save the author to the database
	config.DB.FirstOrCreate(&testAuthor)

	testBook.Title = "Test Book title"
	testBook.Isbn = "ISB-111-111-111"
	testBook.UserID = testUser.ID
	testBook.AuthorID = testAuthor.ID

	// Save the author to the database
	config.DB.FirstOrCreate(&testBook)

	router = routes.SetupRouter()

	// Run tests
	code := m.Run()

	// Cleanup
	config.DB.Delete(&testUser)
	config.DB.Delete(&testBook)
	config.DB.Delete(&testAuthor)

	os.Exit(code)
}
