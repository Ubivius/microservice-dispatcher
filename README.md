# Microservice-Dispatcher
The dispatcher is used for starting game servers in the kubernetes cluster.

When players are logged into their account in the game client they send a request to the disatcher.

Dispatcher calls matchmaking. Once matchmaking returns a slice of players the dispatcher starts the game server.

This is where the kubernetes controller part of this service comes into play.

Once the game server is up and running the dispatcher returns the information for connecting to the new game server to the players.

## Dispatcher endpoints
`GET` `/GameServer` A game client calls this endpoint when a player wants to create a lobby. </br>

Returns
```json
{
    "server_id": "string",
    "server_ip": "string",
    "tcp_port": "int",
    "udp_port": "int",
}
```

`GET` `/IP/{id:[0-9a-z-]+}` A game client calls this endpoint when it wants to join an existing lobby. The "id" is the game id. </br>
```json
{
    "id":    "string, required"
}
```
Returns
```json
{
    "server_id": "string",
    "server_ip": "string",
    "tcp_port": "int",
    "udp_port": "int",
}
```
