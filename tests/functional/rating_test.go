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

type BookRatingRequest struct {
	UserID  uint `json:"user_id"`
	BookID  int
	Rating  int
	Comment string `json:"comment"`
}

func TestGetRatingsRequiresAValidBookID(t *testing.T) {
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
