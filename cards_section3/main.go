package main

import "fmt"

func preMain16() {
	// var card string= "Ace of Spades" is ===
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
