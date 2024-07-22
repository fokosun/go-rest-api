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

type RegisterRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func TestRegisterUserFailsIfFirstnameValidationFails(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Lastname: "test",
		Email:    "user@test.com",
		Password: "examplePass",
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
	assert.Equal(t, "Key: 'User.Firstname' Error:Field validation for 'Firstname' failed on the 'required' tag", errorResponse.Message)
}

func TestRegisterUserFailsIfLastnameValidationFails(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "Fisher",
		Email:     "user@test.com",
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
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "Fisher",
		Lastname:  "Trimii",
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
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "Fisher",
		Lastname:  "Trimii",
		Email:     "invalid",
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
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: testUser.Firstname,
		Lastname:  testUser.Lastname,
		Email:     testUser.Email,
		Password:  "validpassword",
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

func TestRegisterUserFailsIfPasswordValidationFailsIsRequired(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "Fisher",
		Lastname:  "Trimii",
		Email:     "user@test.com",
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
	assert.Equal(t, "Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag", errorResponse.Message)
}

func TestRegisterUserFailsIfPasswordValidationFailsIsLessThanMinLen(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "Fisher",
		Lastname:  "Trimii",
		Email:     "user@test.com",
		Password:  "less",
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
	assert.Equal(t, "Password must be at least 8 characters long.", errorResponse.Message)
}

func TestRegisterUserSucceedsIfEmailDoesNotExistAlready(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "exampleUser",
		Lastname:  "test",
		Email:     "user+3@test.com",
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

	// Find the newly created user by email
	var newUser models.User
	if err := config.DB.Where("email = ?", requestData.Email).First(&newUser).Error; err != nil {
		t.Fatalf("could not find user: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, newUser.Firstname, requestData.Firstname)
	assert.Equal(t, newUser.Lastname, requestData.Lastname)
	assert.Equal(t, newUser.Email, requestData.Email)

	// Delete the newly created user
	t.Cleanup(func() {
		config.DB.Delete(&newUser)
	})
}

func TestGetUsersSucceeds(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var users []models.User
	err = json.Unmarshal(bodyBytes, &users)
	assert.NoError(t, err)

	// Assert that the response is a list
	assert.IsType(t, []models.User{}, users)

	// Assert that each item has the specific keys
	for _, user := range users {
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.Firstname)
		assert.NotEmpty(t, user.Lastname)
		assert.NotEmpty(t, user.Email)
		assert.NotEmpty(t, user.CreatedAt)
		assert.NotEmpty(t, user.UpdatedAt)
	}
}

func TestGetUserByIdRespondsWith404IfUserNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/users/0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))
}

func TestGetUserByIdRespondsWith200IfUserExists(t *testing.T) {
	w := httptest.NewRecorder()

	relativeUrl := "/api/users/" + strconv.Itoa(int(testUser.ID))
	req, _ := http.NewRequest("GET", relativeUrl, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Read the response body
	bodyBytes, err := io.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	// Print the response body for debugging purposes
	fmt.Println(string(bodyBytes))

	// Unmarshal the response body into a slice of users
	var foundUser models.User
	err = json.Unmarshal(bodyBytes, &foundUser)
	assert.NoError(t, err)

	assert.Equal(t, testUser.ID, foundUser.ID)
}

func TestUpdateUserRespondsWith404NotFoundIfUserNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "newFirstname",
		Lastname:  "newLastname",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/api/users/0"
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))

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
	assert.Equal(t, "User not found.", errorResponse.Message)
}

func TestUpdateUserRespondsWith400BadRequestIfUpdatingPasswordAndPasswordFailsValidation(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Password: "bad",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/api/users/" + strconv.Itoa(int(testUser.ID))
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))

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
	assert.Equal(t, "invalid password length", errorResponse.Message)
}

func TestUpdateUserRespondsWith400BadRequestIfTryingToUpdateEmail(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Email: "mybrabdnewemail@test.com",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	baseURL := "http://localhost:8080"
	relativeURL := "/api/users/" + strconv.Itoa(int(testUser.ID))
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))

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
	assert.Equal(t, "email field cannot be updated", errorResponse.Message)
}

func TestCanUpdateUserWithAllowedFields(t *testing.T) {
	w := httptest.NewRecorder()

	requestData := RegisterRequest{
		Firstname: "Freshman",
		Lastname:  "Jamrock",
		Password:  "mynewshinnypassword",
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	var userToUpdate models.User
	userToUpdate.Firstname = "French"
	userToUpdate.Lastname = "Toast"
	userToUpdate.Email = "french.toast@test.com"
	userToUpdate.SetPassword("myValidPassword")
	config.DB.Create(&userToUpdate)

	baseURL := "http://localhost:8080"
	relativeURL := "/api/users/" + strconv.Itoa(int(userToUpdate.ID))
	fullURL := baseURL + relativeURL

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))

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

	var updatedUser models.User
	err = json.Unmarshal(bodyBytes, &updatedUser)
	assert.NoError(t, err)

	assert.Equal(t, updatedUser.Firstname, requestData.Firstname)
	assert.Equal(t, updatedUser.Lastname, requestData.Lastname)

	t.Cleanup(func() {
		config.DB.Delete(&userToUpdate)
	})
}

func TestDeleteUserRespondsWith404NotFoundIfUserNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", "/api/users/0", nil)

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
	assert.Equal(t, "User not found.", errorResponse.Message)
}

func TestDeleteUserSucceedsIfUserFound(t *testing.T) {
	w := httptest.NewRecorder()

	//create a test user
	var newUser models.User
	newUser.Firstname = "test"
	newUser.Lastname = "last"
	newUser.Email = "del@test.com"
	newUser.SetPassword("validpassword")

	// Save the user to the database
	config.DB.Create(&newUser)

	relativeUrl := "/api/users/" + strconv.Itoa(int(newUser.ID))

	req, err := http.NewRequest("DELETE", relativeUrl, nil)

	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	t.Cleanup(func() {
		config.DB.Delete(&newUser)
	})
}
