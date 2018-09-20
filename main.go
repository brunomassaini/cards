package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("All Cards:")
	cards.print()

	// shuffle and distribute card hands
	cards.shuffle()
	hand, remainingCards := deal(cards, 5)

	// print each card separatedly
	fmt.Println("")
	fmt.Println("Player:")
	hand.print()
	fmt.Println("")
	fmt.Println("Dealer:")
	remainingCards.print()

	// write/read it to/from file
	// cards.saveToFile("my_cards")
	// cards = newDeckFromFile("my_cards")
	// cards.print()
}
