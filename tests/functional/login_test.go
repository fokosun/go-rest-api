package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/handlers"
	"github.com/fokosun/go-rest-api/models"
	"github.com/stretchr/testify/assert"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestLoginFailsWhenGivenEmailDoesNotExist(t *testing.T) {
	requestData := LoginRequest{
		Email:    "doesnotexist@test.com",
		Password: "newExamplePass",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/auth/login"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

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
	assert.Equal(t, "invalid email or password", errorResponse.Message)
}

// todo
func TestLoginFailsWhenPasswordCheckFails(t *testing.T) {
	requestData := LoginRequest{
		Email:    testUser.Email,
		Password: "passworddoesnotmatch",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/auth/login"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

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
	assert.Equal(t, "invalid email or password", errorResponse.Message)
}

// to do
func TestLoginSucceedsWhenEmailAndPasswordMatch(t *testing.T) {
	//create new user
	var newUser models.User
	newUser.Firstname = "test"
	newUser.Lastname = "last"
	newUser.Email = "login@test.com"
	newUser.SetPassword("validpassword")

	// Save the user to the database
	config.DB.Create(&newUser)

	requestData := LoginRequest{
		Email:    newUser.Email,
		Password: "validpassword",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/auth/login"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var LoginToken handlers.LoginToken
	err = json.Unmarshal(bodyBytes, &LoginToken)
	assert.NoError(t, err)

	fmt.Printf("LoginToken %v\n", LoginToken)
}
