package main

func main() {
	deck := makeDeck()
	deck.print()

	player1 := makePlayer()
	player1.printInfo()
	player1.addToChipCount(1000)
	player1.printInfo()
	player1.printCatchPhrase()
}
