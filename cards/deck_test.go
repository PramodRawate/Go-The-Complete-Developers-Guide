package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Ace of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndnewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := deck.newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(deck))
	}

	os.Remove("_deckTesting")
}
