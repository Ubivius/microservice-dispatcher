package data

import (
	"github.com/go-playground/validator"
)

func (player *Player) ValidatePlayer() error {
	validate := validator.New()
	return validate.Struct(player)
}
