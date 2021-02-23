package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

// ValidateProduct a product with json validation and IP validator
func (player *Player) ValidateProduct() error {
	validate := validator.New()
	err := validate.RegisterValidation("customip", validateIP)
	if err != nil {
		// Panic if we get this error, that means we are not validating input
		// This will be handled in a better way once we move the JSON validation to accept an interface
		panic(err)
	}

	return validate.Struct(player)
}

// Should validate structure of received IP address
func validateIP(fieldLevel validator.FieldLevel) bool {
	// ip is of format 0.0.0.0
	re := regexp.MustCompile(`[0-9]+.[0-9]+.[0-9]+.[0-9]+`)
	matches := re.FindAllString(fieldLevel.Field().String(), -1)

	return len(matches) == 1
}
