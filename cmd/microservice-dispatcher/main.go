package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ubivius/microservice-dispatcher/pkg/handlers"
	"github.com/Ubivius/pkg-telemetry/metrics"
	"github.com/Ubivius/pkg-telemetry/tracing"
	"github.com/gorilla/mux"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var log = logf.Log.WithName("dispatcher-main")

func main() {
	// Starting k8s logger
	opts := zap.Options{}
	opts.BindFlags(flag.CommandLine)
	newLogger := zap.New(zap.UseFlagOptions(&opts), zap.WriteTo(os.Stdout))
	logf.SetLogger(newLogger.WithName("log"))

	// Starting tracer provider
	tp := tracing.CreateTracerProvider("http://192.168.6.12:14268/api/traces", "microservice-dispatcher-traces")

	// Starting metrics exporter
	metrics.StartPrometheusExporterWithName("dispatcher")

	// Creating handlers
	gameHandler := handlers.NewGameHandler()

	// Mux route handling with gorilla/mux
	router := mux.NewRouter()

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/player", gameHandler.NewPlayer)
	postRouter.Use(gameHandler.MiddlewarePlayerValidation)

	// Get router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/IP/{id:[0-9a-z-]+}", gameHandler.CallGetGameserverIP)
	getRouter.HandleFunc("/GameServer", gameHandler.CallGetReadyGameserver)

	//Health Check
	getRouter.HandleFunc("/health/live", gameHandler.LivenessCheck)
	getRouter.HandleFunc("/health/ready", gameHandler.ReadinessCheck)

	// Server setup
	server := &http.Server{
		Addr:        ":9090",
		Handler:     router,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		log.Info("Starting server", "port", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			log.Error(err, "Server error")
		}
	}()

	// Handle shutdown signals from operating system
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	receivedSignal := <-signalChannel

	log.Info("Received terminate, beginning graceful shutdown", "received_signal", receivedSignal.String())

	// Context cancelling
	timeoutContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cleanly shutdown and flush telemetry on shutdown
	defer func(ctx context.Context) {
		if err := tp.Shutdown(ctx); err != nil {
			log.Error(err, "Error shutting down tracer provider")
		}
	}(timeoutContext)

	// Server shutdown
	_ = server.Shutdown(timeoutContext)
}
