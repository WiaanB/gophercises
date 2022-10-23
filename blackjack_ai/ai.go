package blackjack

import (
	"fmt"
	deck "workspace/deck_of_cards"
)

type AI interface {
	Results(hand [][]deck.Card, dealer []deck.Card)
	Play(hand []deck.Card, dealer deck.Card) Move
	Bet() int
}

type HumanAI struct{}

func (ai *HumanAI) Bet() int {
	return 1
}

func (ai *HumanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	var input string
	for {
		fmt.Printf("Player: %s\nDealer: %s\n%s", hand, dealer, "What will you do? (h)it, (s)tand\n")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return Hit
		case "s":
			return Stand
		default:
			fmt.Println("invalid input: ", input)
		}
	}
}

func (ai *HumanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Printf("==FINAL HANDS==\nPlayer: %s\nScore: %s\nDealer: %s\nScore: %s\n", hand, "pScore", dealer, "dScore")
}

// Filler

type Move func(GameState) GameState

type GameState struct{}

func Hit(gs GameState) GameState {
	return gs
}

func Stand(gs GameState) GameState {
	return gs
}
