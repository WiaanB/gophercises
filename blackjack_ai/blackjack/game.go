package blackjack

import (
	"fmt"
	deck "workspace/deck_of_cards"
)

type state int8

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 100
	}
	if opts.BlackjackPayout == 0.0 {
		opts.BlackjackPayout = 1.5
	}
	g.nDecks = opts.Decks
	g.nHands = opts.Hands
	g.blackjackPayout = opts.BlackjackPayout
	return g
}

type Game struct {
	nDecks          int
	nHands          int
	blackjackPayout float64

	state state
	deck  []deck.Card

	player    []deck.Card
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI dealerAI
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't anyone's turn")
	}
}

func bet(g *Game, ai AI, shuffeled bool) {
	g.playerBet = ai.Bet(shuffeled)
}

func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

func (g *Game) Play(ai AI) int {
	g.deck = deck.New(deck.Deck(g.nDecks), deck.Shuffle())
	min := 52 * g.nDecks / 3
	for i := 0; i < g.nHands; i++ {
		shuffled := false
		if len(g.deck) < min {
			g.deck = deck.New(deck.Deck(g.nDecks), deck.Shuffle())
			shuffled = true
		}
		bet(g, ai, shuffled)
		deal(g)
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}
		EndHand(g, ai)
	}
	return g.balance
}

type Move func(*Game)

func MoveHit(g *Game) {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		MoveStand(g)
	}
}

func MoveStand(g *Game) {
	g.state++
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func EndHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)
	// TODO: figure out winnings and add/substract
	winnings := g.playerBet
	switch {
	case pScore > 21:
		fmt.Println("You busted")
		g.balance--
		winnings *= -1
	case dScore > 21:
		fmt.Println("Dealer busted")
		g.balance++
	case pScore > dScore:
		fmt.Println("You win!")
		g.balance++
	case dScore > pScore:
		fmt.Println("You lose!")
		g.balance--
		winnings *= -1
	case dScore == pScore:
		fmt.Println("Draw!")
		winnings = 0
	}
	g.balance += winnings
	ai.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
	fmt.Println()
}

// Score will take in a hand of cards and return the best blackback score
// with the hand
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			// Ace is worth one, and we make it worth 11
			return minScore + 10
		}
	}
	return minScore
}

// Soft returns true if the score of a hand is a soft score - that is if an ace is being counted as 11 points
func Soft(hand ...deck.Card) bool {
	minscore := minScore(hand...)
	score := Score(hand...)
	return minscore != score
}

func minScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
