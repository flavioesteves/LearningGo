package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	fmt.Println(cards.toString())
}

func main_l24() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}

func main_l25() {
	greeting := "Hi There!"

	fmt.Println([]byte(greeting))
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
