package main

import "fmt"

func main() {
	cards := newDec()
	fmt.Println(cards.toString())
	cards.saveToFile("main_cards")
}
