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

type dealerAI struct{}

func (ai dealerAI) Bet() int {
	//noop
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	//noop
}

func HumanAI() AI {
	return humanAi{}
}

type humanAi struct{}

func (ai humanAi) Bet() int {
	return 1
}

func (ai humanAi) Play(hand []deck.Card, dealer deck.Card) Move {
	var input string
	for {
		fmt.Printf("Player: %s\nDealer: %s\n%s", hand, dealer, "What will you do? (h)it, (s)tand\n")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("invalid input: ", input)
		}
	}
}

func (ai humanAi) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Printf("==FINAL HANDS==\nPlayer: %s\nDealer: %s\n\n", hand, dealer)
}
