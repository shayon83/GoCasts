package main

var deckType string

func main() {
	cards := newDeckFromFile("mycards")
	cards.print()
}
