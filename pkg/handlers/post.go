package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-dispatcher/pkg/data"
)

// NewPlayer creates a new player from the received JSON
func (gameHandler *GameHandler) NewPlayer(responseWriter http.ResponseWriter, request *http.Request) {
	gameHandler.logger.Println("Handle POST for new game request")
	player := request.Context().Value(KeyPlayer{}).(*data.Player)

	data.NewPlayer(player)
	responseWriter.WriteHeader(http.StatusNoContent)
}
