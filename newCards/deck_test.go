package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 36 {
		t.Errorf("Expected a deck of 36, but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected to the first card to be Ace of Spades: %v", d[0])
	}

	if d[len(d) - 1] != "Nine of Clubs" {
		t.Errorf("Expected to the last card to be Nine of Clubs: %v", d[len(d) -1 ])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 36 {
		t.Errorf("Expected 36 Cards in the deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}