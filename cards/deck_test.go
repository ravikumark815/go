package main

import "testing"

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