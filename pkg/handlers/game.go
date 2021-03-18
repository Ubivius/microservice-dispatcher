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

func NewGameHandler(logger *log.Logger) *GameHandler {
	return &GameHandler{logger}
}
