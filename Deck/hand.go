package main

import "fmt"
// Hand is a structure that tracks the current players hand.
type Hand struct {
	cards [6]Card
	points int
}

func (h Hand) print() {
	for i, c := range h.cards {
		if (i == 3) {
			fmt.Println()
		}
		c.print()
		fmt.Printf(" ")
	}
	fmt.Println()
}

