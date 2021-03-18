package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-dispatcher/pkg/data"
)

// MiddlewarePlayerValidation is used to validate incoming player JSONS
func (gameHandler *GameHandler) MiddlewarePlayerValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		player := &data.Player{}

		err := json.NewDecoder(request.Body).Decode(player)
		if err != nil {
			gameHandler.logger.Println("[ERROR] deserializing player", err)
			http.Error(responseWriter, "Error reading player", http.StatusBadRequest)
			return
		}

		// validate the player
		err = player.ValidatePlayer()
		if err != nil {
			gameHandler.logger.Println("[ERROR] validating player", err)
			http.Error(responseWriter, fmt.Sprintf("Error validating player: %s", err), http.StatusBadRequest)
			return
		}

		// Add the player to the context
		ctx := context.WithValue(request.Context(), KeyPlayer{}, player)
		request = request.WithContext(ctx)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(responseWriter, request)
	})
}
