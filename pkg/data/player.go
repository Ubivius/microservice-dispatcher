package data

type Player struct {
	ID int    `json:"id" validate:"required"`
	IP string `json:"ip" validate:"required,ipv4"`
}

// Players is a collection of Player
type Players []*Player

// NewPlayer add the player to the game server queue
func NewPlayer(player *Player) {
	playerList = append(playerList, player)
}

// This will either be a database connection or an in memory slice depending on how we end up handling game servers
var playerList = []*Player{}
