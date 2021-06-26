package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Player - A structure representing a Person Playing the game
type Player struct {
	name        string
	catchPhrase string
	hand        Hand
	roundScore  int
	reader      bufio.Reader
}

// makePlayer - Function for creating a Player
func makePlayer() Player {
	var p Player
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Player Name: ")
	p.name, _ = reader.ReadString('\n')
	p.name = strings.Replace(p.name, "\n", "", -1)
	fmt.Print("Enter Catchphrase: ")
	p.catchPhrase, _ = reader.ReadString('\n')
	p.catchPhrase = strings.Replace(p.catchPhrase, "\n", "", -1)

	return p
}

func (p *Player) addHand(h Hand) {
	p.hand = h
}

// printInfo - Function for printing a players current info
func (p Player) printInfo() {
	fmt.Printf("Player: %v \n", p.name)
}

// printCatchPhrase - Function for Printing a players catchPhrase
func (p Player) printCatchPhrase() {
	fmt.Printf("Player: %v says %v \n", p.name, p.catchPhrase)
}

func (p *Player) playTurn(lastTurn bool, deck *Deck) bool {
	if len(deck.discards) == 0 {
		fmt.Println("Select a move!\n 1. Draw from Deck")
	} else {
		fmt.Printf("Select a move!\n 1. Draw from Deck\n 2. Pick up discard %s\n", deck.peekDiscard())
	}
	var move int
	fmt.Scanln(&move)
	var newCard Card
	if move == 1 {
		newCard = deck.draw()
	} else {
		newCard = deck.drawDiscard()
	}

	fmt.Printf("Select a move with %s!\n 1 to 6 Replace, other discard!\n", newCard.toString())
	p.hand.print()
	fmt.Scanln(&move)
	if move > 0 && move < 7 {
		var oldCard = p.hand.cards[move-1]
		oldCard.turnCardOver()
		p.hand.cards[move-1] = newCard
		fmt.Printf("Replaced %s with %s\n", newCard.toString(), oldCard.toString())
		newCard = oldCard
	}

	deck.discard(newCard)

	if lastTurn || p.hand.allCardsVisable() {
		p.hand.flipHand()
		p.roundScore = p.hand.calculateScore()
	}

	p.hand.print()

	return p.hand.allCardsVisable()
}

func (p Player) getRoundScore() int {
	return p.roundScore
}
