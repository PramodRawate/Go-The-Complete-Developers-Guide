package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDec() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardSuit+" of "+cardValue)
		}
	}

	return cards
}
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:3], d[3:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
