package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-dispatcher/pkg/data"
	"github.com/Ubivius/microservice-dispatcher/pkg/ubiviuscontroller"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
)

func (gameHandler *GameHandler) CallGetGameserverIP(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("dispatcher").Start(request.Context(), "callGetGameserverIP")
	defer span.End()
	log.Info("Calling k8s controller")
	vars := mux.Vars(request)
	id := vars["id"]
	gameServerIP, err := ubiviuscontroller.GetGameserverIP(id)
	if err != nil {
		log.Error(err, "Error while calling controller")
		http.Error(responseWriter, "No ready game server found", http.StatusInternalServerError)
	}

	gameServer := data.GameServer{ID: id, ServerIP: gameServerIP, TCPPort: 9151, UDPPort: 9150}
	err = json.NewEncoder(responseWriter).Encode(gameServer)
	if err != nil {
		log.Error(err, "Error serializing game server")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func (gameHandler *GameHandler) CallGetReadyGameserver(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("dispatcher").Start(request.Context(), "callGetReadyGameserver")
	defer span.End()
	log.Info("Calling k8s controller")

	gameServer, err := ubiviuscontroller.GetReadyGameserver()
	if err != nil {
		log.Error(err, "Error while calling controller")
		http.Error(responseWriter, err.Error(), http.StatusTeapot)
	}

	id := ""
	if len(gameServer.Name) != 0 {
		id = gameServer.Name[len(gameServer.Name)-11:] //only get the game id
	}

	gameServerInfo := data.GameServer{ID: id, ServerIP: gameServer.Status.Address, TCPPort: 9151, UDPPort: 9150}
	err = json.NewEncoder(responseWriter).Encode(gameServerInfo)
	if err != nil {
		log.Error(err, "Error serializing game server")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
