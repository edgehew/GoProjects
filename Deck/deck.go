package main

import "fmt"

// Constant suit values
const (
	SPADE   = "\xE2\x99\xA0"
	CLUB    = "\xE2\x99\xA3"
	HEART   = "\xE2\x99\xA5"
	DIAMOND = "\xE2\x99\xA6"
)

// Deck - Structure extending an array of 52 Cards
type Deck struct {
	cards [52]Card
}

// makeDeck - Function for creating a new deck of playing cards
func makeDeck() Deck {
	var d Deck

	SUITS := [4]string{SPADE, CLUB, HEART, DIAMOND}

	VALUES := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	k := 0
	for _, s := range SUITS {
		for _, v := range VALUES {
			d.cards[k] = makeCard(v, s)
			k++
		}
	}
	return d
}

// print - Function for printing all the cards in the deck
func (d Deck) print() {
	for i, card := range d.cards {
		card.print()
		if i == 12 || i == 25 || i == 38 || i == 51 {
			fmt.Println()
		} else {
			fmt.Print(", ")
		}
	}
}
