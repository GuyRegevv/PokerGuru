package model

import (
	"fmt"
)

// Suit
type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) String() string {
	switch s {
	case Club:
		return "♣"
	case Diamond:
		return "♦"
	case Heart:
		return "♥"
	case Spade:
		return "♠"
	default:
		return "?"
	}
}

// Rank
type Rank int

const (
	RankTwo Rank = iota + 2
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
)

func (r Rank) String() string {
	switch r {
	case RankAce:
		return "A"
	case RankKing:
		return "K"
	case RankQueen:
		return "Q"
	case RankJack:
		return "J"
	case RankTen:
		return "T"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

// Card represents a playing card with a suit and rank.
type Card struct {
	Rank Rank
	Suit Suit
}

type ActionType int

const (
	ActionFold ActionType = iota
	ActionCheck
	ActionCall
	ActionBet
	ActionRaise
)

// Action represents a player's action in the game.
type Action struct {
	PlayerIndex int // index of the player taking the action
	Type        ActionType
	Amount      int // Amount involved in the action, if applicable
}

// Player represents a player in the game.
type Player struct {
	Name  string // Player's name
	Stack int    // Player' s current stack size
	Hand  []Card // Player's current hand
}
