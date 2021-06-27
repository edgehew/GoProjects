package main

import "fmt"

func main() {
	deck1 := makeDeck()
	deck2 := makeDeck()

	deck1.append(deck2)

	deck1.shuffle()

	player1 := makePlayer()
	player1.printInfo()

	player2 := makePlayer()
	player2.printInfo()

	var players []*Player
	players = append(players, &player1)
	players = append(players, &player2)
	scoreboard := makeScoreboard(players)

	for i := 0; i < 9; i++ {
		fmt.Printf("ROUND: %d\n", i+1)
		deck1.shuffle()
		// Deal a hand to each player
		for _, player := range players {
			player.addHand(deck1.deal())
			fmt.Println("----------------------------------------------")
			fmt.Printf("%s's Hand:\n", player.name)
			player.hand.print()
			fmt.Println("----------------------------------------------")
		}

		deck1.discard(deck1.draw())

		// Continue having each player have a turn until all cards get turned over
		var playersDoneCount = 0
		for playersDoneCount < len(players) {
			for _, player := range players {
				fmt.Println("----------------------------------------------")
				var lastTurn = playersDoneCount != 0
				if playersDoneCount < len(players) {
					if player.playTurn(lastTurn, &deck1) {
						playersDoneCount++
					}
				}
			}
		}

		// At end of round calculate scores and add to scoreboard
		for _, player := range players {
			fmt.Println("----------------------------------------------")
			fmt.Printf("%s's Hand:\n", player.name)
			player.hand.print()
			scoreboard.addToScore(player, player.getRoundScore())
			player.resetRoundScore()
			player.discardHand(&deck1)
			fmt.Println("----------------------------------------------")
		}
		playerToEnd := players[0]
		players = players[1:]
		players = append(players, playerToEnd)
	}
}
