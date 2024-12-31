package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNewDeck(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits = 52 cards
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in the new Deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expectedCard := Card{Rank: Ace, Suit: Spade}

	if cards[0] != expectedCard {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expectedCard := Card{Rank: Ace, Suit: Spade}

	if cards[0] != expectedCard {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0

	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Expected 3 Jokers, received: ", count)
	}
}
