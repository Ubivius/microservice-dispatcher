package operator

// Contains information about the game server
type GameServer struct {
	serverIP string
}

func NewGameServer(serverIP string) *GameServer {
	return &GameServer{}
}
