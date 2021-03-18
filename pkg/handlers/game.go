package handlers

import (
	"log"
)

// KeyPlayer is a key used for the Player object inside context
type KeyPlayer struct{}

// GameHandler contains the items common to all game handler functions
type GameHandler struct {
	logger *log.Logger
}

// NewGameHandler returns a pointer to a ProductsHandler with the logger passed as a parameter
func NewGameHandler(logger *log.Logger) *GameHandler {
	return &GameHandler{logger}
}
