package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-dispatcher/pkg/data"
	"github.com/gorilla/mux"
)

func TestValidationMiddlewareWithValidBody(t *testing.T) {
	// Creating request body
	body := &data.Player{
		ID: 1,
		IP: "0.0.0.0",
	}
	bodyBytes, _ := json.Marshal(body)

	request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	productHandler := NewGameHandler(NewTestLogger())

	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/products", productHandler.NewPlayer)
	router.Use(productHandler.MiddlewarePlayerValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestValidationMiddlewareWithInvalidIP(t *testing.T) {
	// Creating request body
	body := &data.Player{
		ID: 1,
		IP: "0.0.0.a",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Error("Body passing to test is not a valid json struct : ", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	productHandler := NewGameHandler(NewTestLogger())

	// Create a router for middleware because linking is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/products", productHandler.NewPlayer)
	router.Use(productHandler.MiddlewarePlayerValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Field validation for 'IP' failed on the 'customip' tag") {
		t.Error("Expected error on field validation for IP but got : ", response.Body.String())
	}
}
