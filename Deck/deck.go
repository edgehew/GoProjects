package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constant suit values
const (
	SPADE   = "\xE2\x99\xA0"
	CLUB    = "\xE2\x99\xA3"
	HEART   = "\xE2\x99\xA5"
	DIAMOND = "\xE2\x99\xA6"
	JOKER   = "\xF0\x9F\x83\x9F"
)

// Deck - Structure extending an array of 52 Cards
type Deck struct {
	cards []Card
}

// makeDeck - Function for creating a new deck of playing cards
func makeDeck() Deck {
	var d Deck

	SUITS := [4]string{SPADE, CLUB, HEART, DIAMOND}
	VALUES := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, s := range SUITS {
		for _, v := range VALUES {
			d.cards = append(d.cards, makeCard(v, s))
		}
	}
	d.cards = append(d.cards, makeCard("Joker", JOKER))
	d.cards = append(d.cards, makeCard("Joker", JOKER))
	return d
}

// print - Function for printing all the cards in the deck
func (d Deck) print() {
	fmt.Println("Printing Deck")
	for i, card := range d.cards {
		card.turnCardOver()
		card.print()
		if (i+1)%12 == 0 && i != 0 {
			fmt.Println()
		} else {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func (d *Deck) append(deck2 Deck) {
	for _, card := range deck2.cards {
		d.cards = append(d.cards, card)
	}
}

func (d *Deck) deal() Hand {
	var delt [6]Card
	copy(delt[:], d.cards[:])
	delt[0].turnCardOver()
	delt[5].turnCardOver()
	d.cards = d.cards[5:]
	return Hand{delt, 0}
}

func (d *Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d.cards {
		newPosition := r.Intn(len(d.cards) - 1)

		d.cards[i], d.cards[newPosition] = d.cards[newPosition], d.cards[i]
	}
}