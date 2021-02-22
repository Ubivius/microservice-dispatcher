package handlers

import (
	"log"
)

// KeyProduct is a key used for the Product object inside context
type KeyProduct struct{}

// GameHandler contains the items common to all product handler functions
type GameHandler struct {
	logger *log.Logger
}

// NewGameHandler returns a pointer to a ProductsHandler with the logger passed as a parameter
func NewGameHandler(logger *log.Logger) *GameHandler {
	return &GameHandler{logger}
}
