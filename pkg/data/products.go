package data

import (
	"fmt"
)

// ErrorProductNotFound : Product specific errors
var ErrorProductNotFound = fmt.Errorf("Product not found")

// Player defines the structure for an API product.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Player struct {
	ID          int     `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Player

// All of these functions will become database calls in the future
// GETTING PRODUCTS

// GetProducts returns the list of products
func GetProducts() Products {
	return playerList
}

// GetProductByID returns a single product with the given id
func GetProductByID(id int) (*Player, error) {
	index := findIndexByProductID(id)
	if index == -1 {
		return nil, ErrorProductNotFound
	}
	return playerList[index], nil
}

// UPDATING PRODUCTS

// UpdateProduct updates the product specified in received JSON
func UpdateProduct(product *Player) error {
	index := findIndexByProductID(product.ID)
	if index == -1 {
		return ErrorProductNotFound
	}
	playerList[index] = product
	return nil
}

// NewPlayer creates a new product
func NewPlayer(product *Player) {
	product.ID = getNextID()
	playerList = append(playerList, product)
}

// DeleteProduct deletes the product with the given id
func DeleteProduct(id int) error {
	index := findIndexByProductID(id)
	if index == -1 {
		return ErrorProductNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	playerList = append(playerList[:index], playerList[index+1])

	return nil
}

// Returns the index of a product in the database
// Returns -1 when no product is found
func findIndexByProductID(id int) int {
	for index, product := range playerList {
		if product.ID == id {
			return index
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our fake database and adds 1
func getNextID() int {
	lastProduct := playerList[len(playerList)-1]
	return lastProduct.ID + 1
}

// playerList is a hard coded list of products for this
// example data source. Should be replaced by database connection
var playerList = []*Player{}
