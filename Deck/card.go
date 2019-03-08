package main

import "fmt"

// Card - Structure representing a Playing card
type Card struct {
	value string
	suit  string
	isFaceUp bool
}

// makeCard - Function for creating new cards
func makeCard(v string, s string) Card {
	return Card{v, s, false}
}

func (c *Card) turnCardOver() {
	c.isFaceUp = true
}

// print - Funciton for printing a card
func (c *Card) print() {
	if (c.isFaceUp) {
		fmt.Print(c.value + c.suit)
	} else {
		fmt.Print("XX")
	}
}
