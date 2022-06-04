package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Player - A structure representing a Person Playing the game
type Player struct {
	name       string
	hand       Hand
	roundScore int
	reader     bufio.Reader
}

// makePlayer - Function for creating a Player
func makePlayer() Player {
	var p Player
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Player Name: ")
	p.name, _ = reader.ReadString('\n')
	p.name = strings.Replace(p.name, "\n", "", -1)

	return p
}

func (p *Player) addHand(h Hand) {
	p.hand = h
}

// printInfo - Function for printing a players current info
func (p Player) printInfo() {
	fmt.Printf("Player: %v \n", p.name)
}

func (p *Player) playTurn(lastTurn bool, deck *Deck) bool {
	if lastTurn {
		fmt.Println("!!! LAST TURN !!!")
	}

	if len(deck.discards) == 0 {
		fmt.Printf("Select a move %s!\n", p.name)
		p.hand.print()
		fmt.Println(" 1. Draw from Deck")
	} else {
		fmt.Printf("Select a move %s!\n", p.name)
		p.hand.print()
		fmt.Printf(" 1. Draw from Deck\n 2. Pick up discard %s\n", deck.peekDiscard())
	}

	var move int
	fmt.Scanln(&move)
	var newCard Card
	if move == 1 {
		newCard = deck.draw()
	} else {
		newCard = deck.drawDiscard()
	}

	fmt.Printf("Select a move with %s!\n Replace [1-6], Discard [ATE]!\n", newCard.toString())
	fmt.Scanln(&move)
	if move > 0 && move < 7 {
		var oldCard = p.hand.cards[move-1]
		oldCard.turnCardOver()
		newCard.turnCardOver()
		p.hand.cards[move-1] = newCard
		fmt.Printf("Replaced %s with %s\n", newCard.toString(), oldCard.toString())
		newCard = oldCard
	}

	deck.discard(newCard)

	if lastTurn || p.hand.allCardsVisible() {
		p.hand.flipHand()
		p.roundScore = p.hand.calculateScore()
	}

	p.hand.print()

	return p.hand.allCardsVisible()
}

func (p Player) getRoundScore() int {
	return p.roundScore
}

func (p *Player) resetRoundScore() {
	p.roundScore = 0
}

func (p *Player) discardHand(d *Deck) {
	for _, c := range p.hand.cards {
		d.discard(c)
	}
}
