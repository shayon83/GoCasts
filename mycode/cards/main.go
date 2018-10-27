package main

var deckType string

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 4)

	hand.print()
	remainingDeck.print()

}
