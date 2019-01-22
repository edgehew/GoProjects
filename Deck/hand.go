package main

// HandClassification classify's a hand.
type HandClassification struct {
	handStrenth int
	cardStrength int
}

// Hand is a structure that tracks the current players hand.
type Hand struct {
	cards          [2]Card
	playbleCards   [3]Card
	classification HandClassification
}

func (h *Hand)classifyHand() HandClassification {
	switch h.classification.handStrenth {
		case 0:
			checkHighCard(h)
		case 1: 
			checkPair(h)
		case 2:
			checkTwoPaid(h)
	}
 
	if Hand.cards[0].value == Hand.cards[1].value {
		currentHand.level = 
	}
}
