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

	cards := deck{"Ace of spades", newCard()}
	cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// Looping through slices in Go
	/* This can be written in receiver functions in deck.go
	for i, card := range cards {
		fmt.Println(i, card)
	}*/
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
