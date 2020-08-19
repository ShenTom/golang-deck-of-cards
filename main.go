package main

import (
	"fmt"
	"github.com/golang-deck-of-cards/deck"
)

func main() {
	d := deck.New(
		deck.NumberOfDecks(1),
		deck.ShuffleDeck(),
		deck.AddJoker(2),
		deck.Filter([]deck.Card{deck.Card{Value: 1, Suit: "hearts"}}),
	)

	fmt.Print(d)
}
