package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// MiddlewarePlayerValidation is used to validate incoming player JSONS
func (gameHandler *GameHandler) MiddlewarePlayerValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		player := &data.Player{}

		err := json.NewDecoder(request.Body).Decode(player)
		if err != nil {
			gameHandler.logger.Println("[ERROR] deserializing product", err)
			http.Error(responseWriter, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = player.ValidateProduct()
		if err != nil {
			gameHandler.logger.Println("[ERROR] validating product", err)
			http.Error(responseWriter, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// Add the product to the context
		ctx := context.WithValue(request.Context(), KeyProduct{}, player)
		request = request.WithContext(ctx)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(responseWriter, request)
	})
}
