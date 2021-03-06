package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-dispatcher/pkg/data"
)

// NewPlayer creates a new player from the received JSON
func (gameHandler *GameHandler) NewPlayer(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("POST request for new game")
	player := request.Context().Value(KeyPlayer{}).(*data.Player)

	data.NewPlayer(player)
	server := data.NewGameServer()
	err := json.NewEncoder(responseWriter).Encode(server)
	if err != nil {
		log.Error(err, "Error creating game server")
		http.Error(responseWriter, "Error creating game server", http.StatusInternalServerError)
	}

	// responseWriter.WriteHeader(http.StatusNoContent)
}
