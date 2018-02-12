package main

import (
	"os"
	"testing"
)

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

func TestNewDeckFromFileSaveToFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	ld := newDeckFromFile("_decktesting")
	if len(ld) != 52 {
		t.Errorf("Expected 52 card god %v", len(ld))
	}

	os.Remove("_decktesting")

}
