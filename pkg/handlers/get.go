package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-dispatcher/pkg/data"
	"github.com/Ubivius/microservice-dispatcher/pkg/ubiviuscontroller"
)

func (gameHandler *GameHandler) CallController(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("Calling k8s controller")
	gameServerIP, err := ubiviuscontroller.GetGameserverIP()
	if err != nil {
		log.Error(err, "Error while calling controller")
	}

	gameServer := data.GameServer{ServerIP: gameServerIP, TCPPort: 9151, UDPPort: 9150}
	err = json.NewEncoder(responseWriter).Encode(gameServer)
	if err != nil {
		log.Error(err, "Error serializing game server")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
