package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/fokosun/go-rest-api/routes"
)

func TestGetBooks(t *testing.T) {
    router := routes.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/books", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "1984")
}

func TestGetBookByID(t *testing.T) {
    router := routes.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/books/1", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "George Orwell")
}
