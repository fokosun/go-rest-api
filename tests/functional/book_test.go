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

type CreateBookRequest struct {
	Title    string `json:"title"`
	Isbn     string `json:"isbn"`
	AuthorID uint
}

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

func TestGetBookReturnsBookIfBookIDExists(t *testing.T) {
	w := httptest.NewRecorder()

	var newUser models.User
	newUser.Firstname = "Tiger"
	newUser.Lastname = "Eisten"
	newUser.Email = "tiger.eisten@test.com"
	newUser.SetPassword("avalidPass")
	config.DB.Create(&newUser)

	var newAuthor models.Author
	newAuthor.Firstname = "Tiger"
	newAuthor.Lastname = "Eisten"
	newAuthor.CreatedBy = newUser.ID
	config.DB.Create(&newAuthor)

	var newBook models.Book
	newBook.Title = "Hansel Un Gretel2"
	newBook.Isbn = "ISN-192-168-71-71"
	newBook.AuthorID = newAuthor.ID

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
	assert.Equal(t, newBook.Title, foundBook.Title)
	assert.Equal(t, newBook.Isbn, foundBook.Isbn)
	assert.Equal(t, newBook.AuthorID, foundBook.Author.ID)
}

func TestCreateBookRequiresBookTitleWhenNull(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := CreateBookRequest{
		Isbn: "ISB-111-111-111",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	fullURL := "http://localhost:8080/books"
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
	assert.Equal(t, "title is required", errorResponse.Message)
}

func TestCreateBookRequiresBookTitleWhenEmpty(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := CreateBookRequest{
		Title: "",
		Isbn:  "ISB-111-111-111",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	fullURL := "http://localhost:8080/books"
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
	assert.Equal(t, "title is required", errorResponse.Message)
}

func TestCreateBookRequiresAuthorID(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := CreateBookRequest{
		Title: "Example Title",
		Isbn:  "ISB-111-111-111",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	fullURL := "http://localhost:8080/books"
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
	assert.Equal(t, "author_id is required", errorResponse.Message)
}

func TestCreateBookRequiresAuthorIDToExist(t *testing.T) {
	w := httptest.NewRecorder()

	var lastRecord models.Author
	config.DB.Last(&lastRecord)

	lastRecordID := lastRecord.ID + 1

	requestData := CreateBookRequest{
		Title:    "Example Title",
		Isbn:     "ISB-111-111-111",
		AuthorID: uint(lastRecordID),
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	fullURL := "http://localhost:8080/books"
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
	assert.Equal(t, "Author not found", errorResponse.Message)
}

func TestCreateBookCanSuccessfullyCreateABookWhenNoErrors(t *testing.T) {
	w := httptest.NewRecorder()

	var lastRecord models.Author
	config.DB.Last(&lastRecord)

	lastRecordID := lastRecord.ID

	requestData := CreateBookRequest{
		Title:    "Example Title",
		Isbn:     "ISB-111-111-111",
		AuthorID: uint(lastRecordID),
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	fullURL := "http://localhost:8080/books"
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
	var foundBook models.Book
	err = json.Unmarshal(bodyBytes, &foundBook)
	assert.NoError(t, err)

	assert.Equal(t, requestData.Title, foundBook.Title)
	assert.Equal(t, requestData.Isbn, foundBook.Isbn)
	assert.Equal(t, requestData.AuthorID, foundBook.Author.ID)
}
