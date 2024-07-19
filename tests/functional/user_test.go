package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/models"
	"github.com/fokosun/go-rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var w *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "pass")
	os.Setenv("DB_NAME", "books_store")
	os.Setenv("DB_PORT", "5432")

	config.ConnectDatabase()
	db := config.DB

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Rating{})

	router = routes.SetupRouter()
	w = httptest.NewRecorder()

	// Run tests
	code := m.Run()

	// Teardown
	// Here you can close connections or clean up resources

	os.Exit(code)
}

func TestRegisterUserFailsIfValidationFails(t *testing.T) {

}

func TestGetUsersSucceeds(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
