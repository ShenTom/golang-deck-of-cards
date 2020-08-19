package deck

import (
	"math/rand"
)

// The Card struct represents a playing card.
type Card struct {
	Value int
	Suit  string
}

//addJoker bool, numDecks int, filterBy int, sort int, sortBy int

// Rule modifies the deck
type Rule func([]Card) []Card

// ShuffleDeck shuffles the deck
func ShuffleDeck() Rule {
	return func(deck []Card) []Card {
		rand.Shuffle(len(deck), func(i, j int) {
			deck[i], deck[j] = deck[j], deck[i]
		})
		return deck
	}
}

// AddJoker adds a number of jokers into the deck
func AddJoker(numOfJokers int) Rule {
	return func(deck []Card) []Card {
		newDeck := deck
		for i := 0; i < numOfJokers; i++ {
			newDeck = append(newDeck, Card{Value: 14, Suit: "Joker"})
		}
		return newDeck
	}
}

// NumberOfDecks determines how many decks to use to construct a new deck
func NumberOfDecks(num int) Rule {
	return func(deck []Card) []Card {
		return createWithNDecks(num)
	}
}

// Filter filters the deck for specfic cards
func Filter(filterBy []Card) Rule {
	return func(deck []Card) []Card {
		return filterDeck(deck, filterBy)
	}
}

func filterDeck(deck []Card, filterBy []Card) []Card {
	var filteredDeck []Card
	for _, element := range deck {
		shouldAdd := true
		for _, filter := range filterBy {
			if filter.Suit == element.Suit && filter.Value == element.Value {
				shouldAdd = false
			}
		}

		if shouldAdd {
			filteredDeck = append(filteredDeck, element)
		}
	}
	return filteredDeck
}

// New creates a deck of cards
// The parameters are executed in order.
func New(params ...Rule) []Card {
	deck := createWithNDecks(1)
	// loop through suits

	for _, element := range params {
		deck = element(deck)
	}
	return deck
}

func createWithNDecks(numOfDecks int) []Card {
	var deck []Card
	suits := [4]string{"spades", "diamonds", "clubs", "hearts"}
	for i := 0; i < numOfDecks; i++ {
		for y := 0; y < 4; y++ {
			for n := 1; n < 14; n++ {
				deck = append(deck, Card{n, suits[y]})
			}
		}
	}

	return deck
}
