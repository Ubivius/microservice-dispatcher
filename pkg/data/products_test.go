package data

import "testing"

func TestChecksValidation(t *testing.T) {
	product := &Player{
		Name:  "Malcolm",
		Price: 2.00,
		SKU:   "abs-abs-abscd",
	}

	err := product.ValidateProduct()

	if err != nil {
		t.Fatal(err)
	}
}
