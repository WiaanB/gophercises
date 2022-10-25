package main

import (
	"fmt"
	"workspace/blackjack_ai/blackjack"
)

func main() {
	game := blackjack.New(blackjack.Options{})
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
