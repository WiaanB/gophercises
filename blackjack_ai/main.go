package main

import (
	"fmt"
	"workspace/blackjack_ai/blackjack"
)

func main() {
	game := blackjack.New(blackjack.Options{
		Decks:           3,
		Hands:           2,
		BlackjackPayout: 1.5,
	})
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
