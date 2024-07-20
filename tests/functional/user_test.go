package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/handlers"
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

func TestRegisterUserFailsIfFirstnameValidationFails(t *testing.T) {
	requestData := RegisterRequest{
		Lastname: "test",
		Email:    "user@test.com",
		Password: "examplePass",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Panic attack: %v\n", err)
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/register"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body
	var errorResponse *handlers.ErrorResponse
	err = json.Unmarshal(bodyBytes, &errorResponse)
	assert.NoError(t, err)

	// Assert that the error message is as expected
	assert.Equal(t, "Key: 'User.Firstname' Error:Field validation for 'Firstname' failed on the 'required' tag", errorResponse.Message)
}

func TestRegisterUserFailsIfLastnameValidationFails(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "Fisher",
		Email:     "user@test.com",
		Password:  "examplePass",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Panic attack: %v\n", err)
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/register"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body
	var errorResponse *handlers.ErrorResponse
	err = json.Unmarshal(bodyBytes, &errorResponse)
	assert.NoError(t, err)

	// Assert that the error message is as expected
	assert.Equal(t, "Key: 'User.Lastname' Error:Field validation for 'Lastname' failed on the 'required' tag", errorResponse.Message)
}

func TestRegisterUserFailsIfEmailValidationFailsNoEmail(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "Fisher",
		Lastname:  "Trimii",
		Password:  "examplePass",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Panic attack: %v\n", err)
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/register"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body
	var errorResponse *handlers.ErrorResponse
	err = json.Unmarshal(bodyBytes, &errorResponse)
	assert.NoError(t, err)

	// Assert that the error message is as expected
	assert.Equal(t, "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag", errorResponse.Message)
}

func TestRegisterUserFailsIfEmailValidationFailsNotValidEmailFormat(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "Fisher",
		Lastname:  "Trimii",
		Email:     "invalid",
		Password:  "examplePass",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Panic attack: %v\n", err)
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/register"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body
	var errorResponse *handlers.ErrorResponse
	err = json.Unmarshal(bodyBytes, &errorResponse)
	assert.NoError(t, err)

	// Assert that the error message is as expected
	assert.Equal(t, "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag", errorResponse.Message)
}

func TestRegisterUserFailsIfEmailValidationFailsNotUnique(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "exampleUser",
		Lastname:  "test",
		Email:     "user@test.com",
		Password:  "newExamplePass",
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

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body
	var errorResponse *handlers.ErrorResponse
	err = json.Unmarshal(bodyBytes, &errorResponse)
	assert.NoError(t, err)

	// Assert that the error message is as expected
	assert.Equal(t, "User already exists.", errorResponse.Message)
}

func TestRegisterUserSucceedsIfEmailDoesNotExistAlready(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "exampleUser",
		Lastname:  "test",
		Email:     "user+1@test.com",
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

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterUserFailsIfPasswordValidationFailsIsRequired(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "Fisher",
		Lastname:  "Trimii",
		Email:     "user@test.com",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Panic attack: %v\n", err)
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/register"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body
	var errorResponse *handlers.ErrorResponse
	err = json.Unmarshal(bodyBytes, &errorResponse)
	assert.NoError(t, err)

	// Assert that the error message is as expected
	assert.Equal(t, "Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag", errorResponse.Message)
}

func TestRegisterUserFailsIfPasswordValidationFailsIsLessThanMinLen(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "Fisher",
		Lastname:  "Trimii",
		Email:     "user@test.com",
		Password:  "less",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Printf("Panic attack: %v\n", err)
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/register"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body
	var errorResponse *handlers.ErrorResponse
	err = json.Unmarshal(bodyBytes, &errorResponse)
	assert.NoError(t, err)

	// Assert that the error message is as expected
	assert.Equal(t, "Password must be at least 8 characters long.", errorResponse.Message)
}

func TestGetUsersSucceeds(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))
}

func TestGetUserByIdRespondsWith404IfUserNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))
}

// todo: revisit this :id 1 may not always exist
func TestGetUserByIdRespondsWith200IfUserExists(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))
}
