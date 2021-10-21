package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-dispatcher/pkg/data"
	"github.com/Ubivius/microservice-dispatcher/pkg/ubiviuscontroller"
	"github.com/gorilla/mux"
)

func (gameHandler *GameHandler) CallGetGameserverIP(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("Calling k8s controller")
	vars := mux.Vars(request)
	id := vars["id"]
	gameServerIP, err := ubiviuscontroller.GetGameserverIP(id)
	if err != nil {
		log.Error(err, "Error while calling controller")
	}

	gameServer := data.GameServer{ID: id, ServerIP: gameServerIP, TCPPort: 9151, UDPPort: 9150}
	err = json.NewEncoder(responseWriter).Encode(gameServer)
	if err != nil {
		log.Error(err, "Error serializing game server")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func (gameHandler *GameHandler) CallGetReadyGameserver(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("Calling k8s controller")

	gameServer, err := ubiviuscontroller.GetReadyGameserver()
	if err != nil {
		log.Error(err, "Error while calling controller")
	}
	id := gameServer.Status.NodeName
	gameServerInfo := data.GameServer{ID: id, ServerIP: gameServer.Status.Address, TCPPort: 9151, UDPPort: 9150}
	err = json.NewEncoder(responseWriter).Encode(gameServerInfo)
	if err != nil {
		log.Error(err, "Error serializing game server")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
