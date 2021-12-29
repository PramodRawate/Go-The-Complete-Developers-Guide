package main

func main() {
	cards := newDec()
	cards.print()

	hand, remainingCards := deal(cards, 3)
	println("********Hand cards")
	hand.print()
	println("********Remaining cards")
	remainingCards.print()

}
