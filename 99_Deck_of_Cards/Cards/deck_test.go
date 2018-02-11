package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck is 52 but got %v cards", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card is not %v but Ace of Hearts", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected last card is no %v but the King of Spades", d[len(d)-1])
	}
}
