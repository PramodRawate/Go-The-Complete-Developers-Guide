package main

import "fmt"

var deckSize int

func main() {
	printState()

	// var card = "Ace of spades"
	card := newCard()
	fmt.Println(card)

	deckSize = 50
	fmt.Println(deckSize)
}

func newCard() string {
	return "Five of Diamonds"
}
