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

func (p *Player) playTurn() bool {
	p.hand.flipHand()
	p.hand.print()
	p.roundScore = p.hand.calculateScore()
	return p.hand.allCardsVisable()
}

func (p Player) getRoundScore() int {
	return p.roundScore
}
