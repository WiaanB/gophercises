package main

import (
	"fmt"
	"workspace/blackjack_ai/blackjack"
	deck "workspace/deck_of_cards"
)

type basicAI struct{}

func (ai *basicAI) Bet(shuffeled bool) int {
	panic("not implemented") // TODO: Implement
}

func (ai *basicAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	panic("not implemented") // TODO: Implement
}

func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	panic("not implemented") // TODO: Implement
}

func main() {
	game := blackjack.New(blackjack.Options{
		Decks:           3,
		Hands:           2,
		BlackjackPayout: 1.5,
	})
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
