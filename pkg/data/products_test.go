package data

import "testing"

func TestChecksValidation(t *testing.T) {
	player := &Player{
		ID: 1,
		IP: "0.0.0.0",
	}

	err := player.ValidatePlayer()

	if err != nil {
		t.Fatal(err)
	}
}
