package main

import "fmt"

// Hand is a structure that tracks the current players hand.
type Hand struct {
	cards  [6]Card
	points int
}

func (h Hand) print() {
	for i, c := range h.cards {
		if i == 3 {
			fmt.Println()
		}
		c.print()
		fmt.Printf(" ")
	}
	fmt.Println()
}

func (h Hand) calculateScore() int {
	score := 0
	for i, c := range h.cards {
		upperCardValue := c.getValue()
		lowerCardValue := h.cards[i+3].getValue()
		if upperCardValue != lowerCardValue {
			score += upperCardValue + lowerCardValue
		}
		if i == 2 {
			break
		}
	}
	score = checkCorners(score, h)
	return score
}

func checkCorners(s int, h Hand) int {
	c := h.cards
	corner1 := c[0].getValue()
	corner2 := c[2].getValue()
	corner3 := c[3].getValue()
	corner4 := c[5].getValue()
	if corner1 == corner2 && corner1 == corner3 && corner1 == corner4 {
		return -10
	}

	return s
}

func (h Hand) allCardsVisable() bool {
	if checkCorners(0, h) == -10 {
		return true
	}

	for _, c := range h.cards {
		if !c.isFaceUp {
			return false
		}
	}
	return true
}

func (h *Hand) flipHand() {
	for i := range h.cards {
		h.cards[i].turnCardOver()
	}
}
