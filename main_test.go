package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/add/{value1}/{value2}", AddHandler)
	return router
}

func TestAddHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/add/2/3", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, "4", response.Body.String(), "Expect response '5'")
}
