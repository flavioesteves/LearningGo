package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}

func preMainUntil15() {
	// var card string= "Ace of Spades" is ===
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}

func main_l16() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
