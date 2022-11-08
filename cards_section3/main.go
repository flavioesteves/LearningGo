package main

import "fmt"

func main() {
	cards := deck{newCard(), "Ace of Diamonds"} // deck.go
	cards = append(cards, "Six of Spades")      // return a new slice

	cards.print()
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
