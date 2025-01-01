package deck

import (
	"fmt"
	"math/rand"
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

func TestShuffle(t *testing.T) {
	// on Go Playground, rand.NewSource(0) will give a specific result (lol that's why we used timestamps there to get different values)
	// First call of shuffleRand.Perm(52) should be: [40, 35, ...]
	// Lets use that

	shuffleRand = rand.New(rand.NewSource(0))

	originalDeck := New()
	firstCard := originalDeck[40]
	secondCard := originalDeck[35]
	cards := New(Shuffle)

	if cards[0] != firstCard {
		t.Errorf("Expected the first card to be %s, received %s.", firstCard, cards[0])
	}

	if cards[1] != secondCard {
		t.Errorf("Expected the first card to be %s, received %s.", secondCard, cards[1])
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

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))

	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))

	expected := 13 * 4 * 3
	if len(cards) != expected {
		t.Errorf("Expected %d cards, received %d cards", expected, len(cards))
	}
}
