package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/fokosun/go-rest-api/handlers"
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

// to do
func TestLoginSucceedsWhenEmailAndPasswordMatch(t *testing.T) {

}
