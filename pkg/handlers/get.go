package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-dispatcher/pkg/ubiviuscontroller"
)

func (gameHandler *GameHandler) CallController(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("Calling k8s controller")
	err := ubiviuscontroller.CreateGameserver()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
	responseWriter.WriteHeader(http.StatusOK)
}
