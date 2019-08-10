package main

func main() {
	deck1 := makeDeck()
	deck2 := makeDeck()

	deck1.append(deck2)

	deck1.shuffle()
	deck1.print()

	player1 := makePlayer()
	player1.printInfo()

	var players []Player
	players = append(players, player1)
	scoreboard := makeScoreboard(players)
	player1.addHand(deck1.deal())
	if player1.playTurn() {
		scoreboard.addToScore(&player1, player1.getRoundScore())
	}

}