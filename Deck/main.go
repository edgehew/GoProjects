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

	var players []Player
	players = append(players, player1)
	scoreboard := makeScoreboard(players)
	hand := deck1.deal()
	hand.print()
	scoreboard.addToScore(&player1, -10)
}
