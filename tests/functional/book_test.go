package tests

import (
	"database/sql"
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

func TestGetBooksSucceeds(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var books []models.Book
	err = json.Unmarshal(bodyBytes, &books)
	assert.NoError(t, err)

	// Assert that the response is a list
	assert.IsType(t, []models.Book{}, books)

	// Assert that each item has the specific keys
	for _, book := range books {
		assert.NotEmpty(t, book.ID)
		assert.NotEmpty(t, book.Title)
		assert.NotEmpty(t, book.Isbn)
		assert.NotEmpty(t, book.CreatedAt)
		assert.NotEmpty(t, book.UpdatedAt)
	}
}

func TestGetBooksRespondsWith404NotFoundWhenBookIDDoesNotExist(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/books/0", nil)
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
	assert.Equal(t, "Book not found", errorResponse.Message)
}

func TestGetBooksReturnsBookIfBookIDExists(t *testing.T) {
	w := httptest.NewRecorder()

	var newBook models.Book
	newBook.Title = "Hansel Un Gretel"
	newBook.Isbn = "ISN-192-168-71-71"
	authorID := sql.NullInt64{Int64: int64(testUser.ID), Valid: true}
	newBook.AuthorID = authorID

	config.DB.Create(&newBook)

	relativeUrl := "/books/" + strconv.Itoa(int(newBook.ID))

	req, _ := http.NewRequest("GET", relativeUrl, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var foundBook models.Book
	err = json.Unmarshal(bodyBytes, &foundBook)
	assert.NoError(t, err)

	assert.Equal(t, newBook.ID, foundBook.ID)
}
