package router

import (
	"net/http"

	"github.com/Ubivius/microservice-dispatcher/pkg/handlers"
	"github.com/Ubivius/pkg-telemetry/metrics"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

// Mux route handling with gorilla/mux
func New(gameHandler *handlers.GameHandler) *mux.Router {
	log.Info("Starting router")
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("dispatcher"))
	router.Use(metrics.RequestCountMiddleware)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/player", gameHandler.NewPlayer)
	postRouter.Use(gameHandler.MiddlewarePlayerValidation)

	// Get router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/IP/{id:[0-9a-z-]+}", gameHandler.CallGetGameserverIP)
	getRouter.HandleFunc("/GameServer", gameHandler.CallGetReadyGameserver)

	// Health Check
	getRouter.HandleFunc("/health/live", gameHandler.LivenessCheck)
	getRouter.HandleFunc("/health/ready", gameHandler.ReadinessCheck)
	return router
}
