package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var w *httptest.ResponseRecorder

type RegisterRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "pass")
	os.Setenv("DB_NAME", "books_store")
	os.Setenv("DB_PORT", "5432")

	config.ConnectDatabase()

	router = routes.SetupRouter()
	w = httptest.NewRecorder()

	// Run tests
	code := m.Run()

	// Teardown
	// Here you can close connections or clean up resources

	os.Exit(code)
}

func TestRegisterUserFailsIfValidationFails(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "exampleUser",
		Lastname:  "test",
		Email:     "user+1@example.com",
		Password:  "examplePass",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/register"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Use http.Client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUsersSucceeds(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
