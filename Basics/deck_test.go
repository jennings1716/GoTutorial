package main

import "testing"

func TestNewDeck(t *testing.T) {
	tescards := deck{}
	tescards = tescards.addCard("card1")
	if len(tescards) != 1 {
		t.Errorf("Add card functionality is not working Expected: 1 But %v", len(tescards))
	}
}
