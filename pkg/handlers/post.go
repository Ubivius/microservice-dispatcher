package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// NewPlayer creates a new product from the received JSON
func (productHandler *GameHandler) NewPlayer(responseWriter http.ResponseWriter, request *http.Request) {
	productHandler.logger.Println("Handle POST for new game request")
	product := request.Context().Value(KeyProduct{}).(*data.Player)

	data.NewPlayer(product)
	responseWriter.WriteHeader(http.StatusNoContent)
}
