package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Hand struct{}

// Player - A structure representing a Person Playing the game
type Player struct {
	name        string
	catchPhrase string
	chipCount   int
	hand        Hand
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
	p.chipCount = 10000

	return p
}

// addToChipCount - Function for adding to a players chip count
func (p *Player) addToChipCount(amount int) {
	p.chipCount += amount
}

// printInfo - Function for printing a players current info
func (p Player) printInfo() {
	fmt.Printf("Player: %v has $%v \n", p.name, p.chipCount)
}

// printCatchPhrase - Function for Printing a players catchPhrase
func (p Player) printCatchPhrase() {
	fmt.Printf("Player: %v says %v \n", p.name, p.catchPhrase)
}
