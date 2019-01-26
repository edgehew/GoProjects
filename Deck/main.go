package main

func main() {
	deck1 := makeDeck()
	deck2 := makeDeck()

	for _, card := range deck2.cards {
		deck1.cards = append(deck1.cards, card)
	}
	deck1.print()

	player1 := makePlayer()
	player1.printInfo()
	player1.addToPoints(1000)
	player1.printInfo()
	player1.printCatchPhrase()
}
