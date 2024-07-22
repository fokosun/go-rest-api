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

	"github.com/bxcodec/faker/v3"
	"github.com/fokosun/go-rest-api/config"
	"github.com/fokosun/go-rest-api/handlers"
	"github.com/fokosun/go-rest-api/models"
	"github.com/stretchr/testify/assert"
)

type BookRatingRequest struct {
	UserID  uint `json:"user_id"`
	BookID  int
	Rating  int
	Comment string `json:"comment"`
}

func TestCreateOrUpdateRatingRequiresAValidBookID(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := BookRatingRequest{
		Rating:  5,
		Comment: "Nice read!",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	var lastBookID uint

	var lastBook models.Book
	if err := config.DB.Last(&lastBook).Error; err != nil {
		lastBookID = testBook.ID + uint(1)
	} else {
		lastBookID = lastBook.ID + uint(1)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/api/books/" + strconv.Itoa(int(lastBookID)) + "/ratings"
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
	assert.Equal(t, "Book not found", errorResponse.Message)
}

func TestCreateOrUpdateWillCreateANewRatingIfRatingNotExisting(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := BookRatingRequest{
		Rating:  5,
		Comment: "Nice read!",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/api/books/" + strconv.Itoa(int(testBook.ID)) + "/ratings"
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
	var newBookRating models.Rating
	err = json.Unmarshal(bodyBytes, &newBookRating)
	assert.NoError(t, err)

	t.Cleanup(func() {
		config.DB.Delete(&newBookRating)
	})
}

func TestCreateOrUpdateWillUpdateTheExistingBookRating(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := BookRatingRequest{
		Rating:  5,
		Comment: "My New Comment On this book!",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	var existingBookRating models.Rating
	existingBookRating.UserID = testUser.ID
	existingBookRating.BookID = int(testBook.ID)
	existingBookRating.Comment = "A Nice Read!"
	config.DB.Create(&existingBookRating)

	baseURL := "http://localhost:8080"
	relativeURL := "/api/books/" + strconv.Itoa(int(existingBookRating.BookID)) + "/ratings"
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

	var updatedBookRating models.Rating
	err = json.Unmarshal(bodyBytes, &updatedBookRating)
	assert.NoError(t, err)

	assert.Equal(t, updatedBookRating.Comment, "My New Comment On this book!")

	t.Cleanup(func() {
		config.DB.Delete(&existingBookRating)
	})
}

func TestGetRatingByBookIDCanReturnTheRatingsForAGivenBookID(t *testing.T) {
	w := httptest.NewRecorder()

	relativeUrl := "/api/books/" + strconv.Itoa(int(testBook.ID)) + "/ratings"
	req, _ := http.NewRequest("GET", relativeUrl, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))
}

func TestGetRatingsCanReturnAllBookRatings(t *testing.T) {
	w := httptest.NewRecorder()

	//Create ratigs for the same BookID but diff users
	for i := 0; i < 5; i++ {
		var newUser models.User
		newUser.Firstname = faker.FirstNameFemale()
		newUser.Lastname = faker.LastName()
		newUser.Email = faker.Email() + "@test.com"
		newUser.SetPassword("myValidPassword")
		config.DB.Create(&newUser)

		var bookRating models.Rating
		bookRating.BookID = int(testBook.ID)
		bookRating.UserID = newUser.ID
		bookRating.Comment = "Nide book"
		config.DB.Create(&bookRating)
	}

	req, _ := http.NewRequest("GET", "/api/books/ratings", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var ratings []models.Rating
	err = json.Unmarshal(bodyBytes, &ratings)
	assert.NoError(t, err)

	// Assert that the response is a list
	assert.IsType(t, []models.Rating{}, ratings)

	// Assert that each item has the specific keys
	for _, rating := range ratings {
		assert.NotEmpty(t, rating.ID)
		assert.NotEmpty(t, rating.Rating)
		assert.NotEmpty(t, rating.Comment)
		assert.NotEmpty(t, rating.BookID)
	}
}
