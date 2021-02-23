package data

// Player defines the structure for an API product.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Player struct {
	ID          int     `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	IP          string  `json:"ip" validate:"required,ip"`
}

// Products is a collection of Product
type Products []*Player

// NewPlayer add the player to the game server queue
func NewPlayer(player *Player) {
	playerList = append(playerList, player)
}

// This will either be a database connection or an in memory array depending on how we end up handling game servers
var playerList = []*Player{}
