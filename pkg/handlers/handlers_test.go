package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ubivius/microservice-dispatcher/pkg/data"
)

func TestNewPlayer(t *testing.T) {
	// Creating request body
	body := &data.Player{
		ID: 1,
		IP: "0.0.0.0",
	}

	request := httptest.NewRequest(http.MethodPost, "/player", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyPlayer{}, body)
	request = request.WithContext(ctx)

	gameHandler := NewGameHandler()
	gameHandler.NewPlayer(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}
