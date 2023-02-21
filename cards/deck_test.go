package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(d))
	}

	if d[0] != "A-Diamonds" {
		t.Errorf("Expected A-Diamonds as first card, got %v", d[0])
	}

	if d[len(d)-1] != "K-Clubs" {
		t.Errorf("Expected K-Clubs as last card, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if (len(loadedDeck) != 52) {
		t.Errorf("Expected 52 cards in new deck from file, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}