package main

// HandClassification classify's a hand.
type HandClassification struct {
	handStrenth  int
	cardStrength int
}

// Hand is a structure that tracks the current players hand.
type Hand struct {
	cards          [2]Card
	playbleCards   [3]Card
	classification HandClassification
}
