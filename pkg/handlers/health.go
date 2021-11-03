package handlers

import (
	"net/http"
)

// LivenessCheck determine when the application needs to be restarted
func (gameHandler *GameHandler) LivenessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
}

//ReadinessCheck verifies that the application is ready to accept requests
func (gameHandler *GameHandler) ReadinessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
}
