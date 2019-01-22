package main

import "fmt"

// Card - Structure representing a Playing card
type Card struct {
	value string
	suit  string
}

// makeCard - Function for creating new cards
func makeCard(v string, s string) Card {
	return Card{v, s}
}

// print - Funciton for printing a card
func (c Card) print() {
	fmt.Print(c.value + c.suit)
}
