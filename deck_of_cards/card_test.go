package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Print(Card{Rank: Ace, Suit: Heart})
	fmt.Print(Card{Rank: Two, Suit: Spade})
	fmt.Print(Card{Rank: Nine, Suit: Diamond})
	fmt.Print(Card{Rank: Jack, Suit: Club})
	fmt.Print(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("wrong number of cards, expecting 52, and got ", len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("expected Ace of Spaces as first card. received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("expected Ace of Spaces as first card. received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("expected 3 jokers, received:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("expected all 2s and 3s to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Errorf("expected %d cards, recevied %d cards", 13*4*3, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// first call to shuffleRand.Perm(52) should be
	// [40 35 50 0 44 7 1 16 13 4 21 12 23 34 19 11 42 20 17 48 27 9 43 46 47 45 5 49 51 30 41 26 25 32 39 28 37 31 33 10 22 8 6 29 36 18 14 2 15 3 38 24]
	shuffleRand = rand.New(rand.NewSource(0))
	original := New()
	first := original[40]
	second := original[35]
	cards := New(Shuffle(original))
	if cards[0] != first {
		t.Errorf("expected the first card to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("expected the first card to be %s, received %s", second, cards[1])
	}
}
