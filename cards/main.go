package main

import "fmt"

var deckSize int

func main() {
	// printState()

	deckSize = 50
	fmt.Println(deckSize)

	deckSize = 50
	fmt.Println(deckSize)

	// var card = "Ace of spades"
	card := newCard()
	fmt.Println(card)

	cards := []string{"Ace of spades", newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
