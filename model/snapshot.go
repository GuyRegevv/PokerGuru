package model

type Round string

const (
	RoundPreFlop Round = "PreFlop"
	RoundFlop    Round = "Flop"
	RoundTurn    Round = "Turn"
	RoundRiver   Round = "River"
)

type PlayerSnapshot struct {
	Name       string
	Stack      int
	IsHero     bool
	HasFolded  bool
	IsAllIn    bool
	CurrentBet int // Amount in front of player this street
	Cards      [2]Card
}

// The "Single Source of Truth" for the Advisor
type GameSnapshot struct {
	Round          Round
	Pot            int
	CommunityCards []Card

	Players    []PlayerSnapshot
	HeroIndex  int
	ToActIndex int

	CallAmount int
	MinRaise   int
	LastAction *Action
}
