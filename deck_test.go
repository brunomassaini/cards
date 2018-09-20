package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	fmt.Println("Testing newDeck")

	d := newDeck()

	// Testing variables
	length := 16
	firstCard := "Ace of Spades"
	lastCard := "Four of Clubs"
	result := ""

	result = "PASS"
	if len(d) != length {
		result = "FAILED"
		t.Errorf("Expected deck length of %v but got %v", length, len(d))
	}
	fmt.Println("(", result, ") 1. Test length")

	result = "PASS"
	if d[0] != firstCard {
		result = "FAILED"
		t.Errorf("Expected first card %v but got %v", firstCard, d[0])
	}
	fmt.Println("(", result, ") 2. Test First Card")

	result = "PASS"
	if d[len(d)-1] != lastCard {
		result = "FAILED"
		t.Errorf("Expected last card %v but got %v", lastCard, d[0])
	}
	fmt.Println("(", result, ") 3. Test Last Card")

	fmt.Println("")
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fmt.Println("Testing save/read to/from file")

	os.Remove("_decktesting")

	// Testing variables
	length := 16
	firstCard := "Ace of Spades"
	lastCard := "Four of Clubs"
	result := ""

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	result = "PASS"
	if len(loadedDeck) != length {
		result = "FAILED"
		t.Errorf("Expected deck length of %v but got %v", length, len(loadedDeck))
	}
	fmt.Println("(", result, ") 1. Test length")

	result = "PASS"
	if loadedDeck[0] != firstCard {
		result = "FAILED"
		t.Errorf("Expected first card %v but got %v", firstCard, loadedDeck[0])
	}
	fmt.Println("(", result, ") 2. Test First Card")

	result = "PASS"
	if loadedDeck[len(d)-1] != lastCard {
		result = "FAILED"
		t.Errorf("Expected last card %v but got %v", lastCard, loadedDeck[0])
	}
	fmt.Println("(", result, ") 3. Test Last Card")

	os.Remove("_decktesting")
	fmt.Println("")
}
