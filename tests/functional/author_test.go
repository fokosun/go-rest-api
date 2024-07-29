package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/handlers"
	"github.com/fokosun/go-rest-api/models"
	"github.com/stretchr/testify/assert"
)

type CreateAuthorRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gravatar  string `json:"gravatar"`
}

func TestAuthorFirstnameIsRequiredExpect400Badrequest(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Lastname: "Trimii",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/api/users/authors"
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
	assert.Equal(t, "Key: 'Author.Firstname' Error:Field validation for 'Firstname' failed on the 'required' tag", errorResponse.Message)
}

func TestAuthorLastnameIsRequiredExpect400Badrequest(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "Trimii",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/api/users/authors"
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
	assert.Equal(t, "Key: 'Author.Lastname' Error:Field validation for 'Lastname' failed on the 'required' tag", errorResponse.Message)
}

func TestExistingUserCanSuccessfullyCreateAuthor(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "first",
		Lastname:  "last",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/api/users/authors"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var foundAuthor models.Author
	err = json.Unmarshal(bodyBytes, &foundAuthor)
	assert.NoError(t, err)

	t.Cleanup(func() {
		config.DB.Delete(&foundAuthor)
	})
}

func TestGetAuthorsSucceeds(t *testing.T) {
	w := httptest.NewRecorder()

	relativeURL := "/api/users/authors"

	req, _ := http.NewRequest("GET", relativeURL, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var authors []models.Author
	err = json.Unmarshal(bodyBytes, &authors)
	assert.NoError(t, err)

	// Assert that the response is a list
	assert.IsType(t, []models.Author{}, authors)

	// Assert that each item has the specific keys
	for _, user := range authors {
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.Firstname)
		assert.NotEmpty(t, user.Lastname)
		assert.NotEmpty(t, user.CreatedAt)
		assert.NotEmpty(t, user.UpdatedAt)
	}
}

func TestGetAuthorByIdRespondsWith404IfUserNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	relativeURL := "/api/users/authors/0"

	req, _ := http.NewRequest("GET", relativeURL, nil)
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
	assert.Equal(t, "Author not found", errorResponse.Message)
}

func TestGetAuthorByIdRespondsWith200IfUserExists(t *testing.T) {
	w := httptest.NewRecorder()

	var newAuthor models.Author
	newAuthor.Firstname = "Elmo"
	newAuthor.Lastname = "Fenmo"

	config.DB.Create(&newAuthor)

	relativeUrl := "/api/users/authors/" + strconv.Itoa(int(newAuthor.ID))

	req, _ := http.NewRequest("GET", relativeUrl, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var foundAuthor models.Author
	err = json.Unmarshal(bodyBytes, &foundAuthor)
	assert.NoError(t, err)

	assert.Equal(t, newAuthor.ID, foundAuthor.ID)

	t.Cleanup(func() {
		config.DB.Delete(&newAuthor)
	})
}
