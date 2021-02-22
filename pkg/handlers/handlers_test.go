package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// Move to util package in Sprint 9, should be a testing specific logger
func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestAddProduct(t *testing.T) {
	// Creating request body
	body := &data.Product{
		Name:        "addName",
		Description: "addDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}

	request := httptest.NewRequest(http.MethodPost, "/products", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyProduct{}, body)
	request = request.WithContext(ctx)

	productHandler := NewProductsHandler(NewTestLogger())
	productHandler.AddProduct(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}
