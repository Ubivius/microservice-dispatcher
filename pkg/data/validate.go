package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

// ValidateProduct a product with json validation and IP validator
func (product *Player) ValidateProduct() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", validateIP)
	if err != nil {
		// Panic if we get this error, that means we are not validating input
		// This will be handled in a better way once we move the JSON validation to accept an interface
		panic(err)
	}

	return validate.Struct(product)
}

// Should validate structure of received IP address
func validateIP(fieldLevel validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fieldLevel.Field().String(), -1)

	return len(matches) == 1
}
