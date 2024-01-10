package main

import (
	"os"
	"testing"
)

// Deciding What to test

// newDeck
func TestNewDeck(t *testing.T) {
	// Arrange
	d := newDeck()
	deckSize := 52
	firstCard := "Ace of Spades"
	lastCard := "King of Clubs"

	// Act
	// Assert
	if len(d) != deckSize {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Expected first card of Ace of Spades, but got  %v", d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("Expected first card of King of Clubs, but got  %v", d[len(d)-1])
	}

}

// saveToFile _decktesting
// newDeckFromFile
func TestSaveToDeckHandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52m but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
