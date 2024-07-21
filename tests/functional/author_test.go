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

type CreateAuthorRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gravatar  string `json:"gravatar"`
}

func TestAuthorFirstnameIsRequiredExpect404Badrequest(t *testing.T) {
	requestData := RegisterRequest{
		Lastname: "Trimii",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/users/authors"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

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
	assert.Equal(t, "Key: 'Author.Firstname' Error:Field validation for 'Firstname' failed on the 'required' tag", errorResponse.Message)
}

func TestAuthorLastnameIsRequiredExpect404Badrequest(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "Trimii",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/users/authors"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

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
	assert.Equal(t, "Key: 'Author.Lastname' Error:Field validation for 'Lastname' failed on the 'required' tag", errorResponse.Message)
}

func TestExistingUserCanSuccessfullyCreateAuthor(t *testing.T) {
	requestData := RegisterRequest{
		Firstname: "first",
		Lastname:  "last",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/users/authors"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
