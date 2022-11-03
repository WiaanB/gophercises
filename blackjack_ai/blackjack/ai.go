package blackjack

import (
	"fmt"
	deck "workspace/deck_of_cards"
)

type AI interface {
	Bet(shuffeled bool) int
	Results(hand [][]deck.Card, dealer []deck.Card)
	Play(hand []deck.Card, dealer deck.Card) Move
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffeled int) int {
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

func (ai humanAi) Bet(shuffeled bool) int {
	if shuffeled {
		fmt.Println("The deck was just shuffeled")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAi) Play(hand []deck.Card, dealer deck.Card) Move {
	var input string
	for {
		fmt.Printf("Player: %s\nDealer: %s\n%s", hand, dealer, "What will you do? (h)it, (s)tand, (d)ouble\n")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		default:
			fmt.Println("invalid input: ", input)
		}
	}
}

func (ai humanAi) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Printf("==FINAL HANDS==\nPlayer: %s\nDealer: %s\n\n", hand, dealer)
}
