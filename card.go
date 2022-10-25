package main

type Card struct {
	Suit   int `json: "suit"`
	Value  int `json: "value"`
	hidden bool
	set    bool
}

func drawCard(cardPos *int, hand *[]Card, deck *[]Card) {
	if *cardPos == -2 {
		*cardPos = -1
	}

	if len(*deck) > 0 {
		*cardPos += 1
		moveCard(hand, deck)
	} else {
		*deck = getDeck()
		*deck = shuffleDeck(*deck)
		*cardPos = 0
		moveCard(hand, deck)
	}
}

func moveCard(hand *[]Card, deck *[]Card) {
	*hand = append(*hand, (*deck)[0])
	(*hand)[len(*hand)-1].set = true
	*deck = (*deck)[1:]
}