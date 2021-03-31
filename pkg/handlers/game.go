package handlers

// KeyPlayer is a key used for the Player object inside context
type KeyPlayer struct{}

// GameHandler contains the items common to all game handler functions
type GameHandler struct {
}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}
