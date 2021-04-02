package data

// GameServer contains the information returned to the game client once the game server is created
type GameServer struct {
	ServerIP string `json:"server_ip"`
	TCPPort  int    `json:"tcp_port"`
	UDPPort  int    `json:"udp_port"`
}

// This function currently returns the values that we need for local testing with the current version of the client-server setup
func NewGameServer() *GameServer {
	return &GameServer{ServerIP: "192.168.0.141", TCPPort: 9051, UDPPort: 9050}
}
