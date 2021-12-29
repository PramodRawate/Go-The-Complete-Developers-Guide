package main

func main() {
	cards := newDec()
	cards.saveToFile("main_cards")

	cardsFromFile := cards.newDeckFromFile("./main_cards")
	cardsFromFile.print()
}
