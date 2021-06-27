package main

import (
	"fmt"
	"log"
	"strconv"
)

// Card - Structure representing a Playing card
type Card struct {
	value    string
	suit     string
	isFaceUp bool
}

// makeCard - Function for creating new cards
func makeCard(v string, s string) Card {
	return Card{v, s, false}
}

func (c *Card) turnCardOver() {
	c.isFaceUp = true
}

func (c *Card) turnCardDown() {
	c.isFaceUp = false
}

func (c *Card) toString() string {
	return c.value + c.suit
}

// print - Funciton for printing a card
func (c *Card) print() {
	if c.isFaceUp {
		fmt.Print(c.toString())
	} else {
		fmt.Print("XX")
	}
}

func (c *Card) getValue() int {
	if c.isFaceUp {
		return translateValue(c.value)
	}
	return 0
}

func translateValue(s string) int {
	// Special Values that do not match card number
	if s == "2" {
		return -2
	} else if s == "Joker" {
		return -5
	} else if s == "A" {
		return 1
	} else if s == "K" {
		return 0
	} else if s == "Q" || s == "J" {
		return 10
	} else {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return i
	}
}
