package engine

import (
	"fmt"
	"poker-guru/model"
)

// Game represents the state of a poker game.
type Game struct {
	Deck           *Deck
	Players        []model.Player
	CommunityCards []model.Card

	Pot         int
	DealerIndex int
	SmallBlind  int
	BigBlind    int
}

// // NewGame initializes a new game with players.
func NewGame(players []model.Player) *Game {
	fmt.Println("Starting a new game...")
	return &Game{
		Deck:           NewDeck(),
		Players:        players,
		CommunityCards: []model.Card{},
		Pot:            0,
		DealerIndex:    0,
		SmallBlind:     5,
		BigBlind:       10,
	}
}

func (g *Game) StartHand() error {
	fmt.Println("Starting a new Hand...")

	g.Deck = NewDeck()

	g.CommunityCards = nil

	for i := range g.Players {
		g.Players[i].Hand = nil
	}

	if len(g.Players) == 0 {
		return fmt.Errorf("no players in game")
	}
	g.DealerIndex = (g.DealerIndex + 1) % len(g.Players)

	g.Pot = 0

	g.postBlinds()

	return nil
}

func (g *Game) DealHoleCards() error {
	fmt.Println("Dealing Cards...")
	for i := range g.Players {
		card1, err := g.Deck.DrawCard()
		if err != nil {
			return err
		}
		card2, err := g.Deck.DrawCard()
		if err != nil {
			return err
		}

		// assign a new slice with two cards
		g.Players[i].Hand = []model.Card{card1, card2}
	}
	return nil
}

func (g *Game) DealFlop() error {
	fmt.Println("Dealing Flop...")
	for i := 0; i < 3; i++ {
		card, err := g.Deck.DrawCard()
		if err != nil {
			return err
		}
		g.CommunityCards = append(g.CommunityCards, card)
	}
	return nil
}

func (g *Game) DealTurn() error {
	fmt.Println("Dealing turn...")
	card, err := g.Deck.DrawCard()
	if err != nil {
		return err
	}
	g.CommunityCards = append(g.CommunityCards, card)
	return nil
}

func (g *Game) DealRiver() error {
	fmt.Println("Dealing river...")
	card, err := g.Deck.DrawCard()
	if err != nil {
		return err
	}
	g.CommunityCards = append(g.CommunityCards, card)
	return nil
}

func (g *Game) postBlinds() {
	if len(g.Players) < 2 {
		return // can't really play, but don't crash
	}

	// small blind is next to dealer
	sbIndex := (g.DealerIndex + 1) % len(g.Players)
	bbIndex := (g.DealerIndex + 2) % len(g.Players)

	// small blind
	sbAmount := g.SmallBlind
	if g.Players[sbIndex].Stack < sbAmount {
		sbAmount = g.Players[sbIndex].Stack // all-in if too small
	}
	g.Players[sbIndex].Stack -= sbAmount
	g.Pot += sbAmount

	// big blind
	bbAmount := g.BigBlind
	if g.Players[bbIndex].Stack < bbAmount {
		bbAmount = g.Players[bbIndex].Stack
	}
	g.Players[bbIndex].Stack -= bbAmount
	g.Pot += bbAmount

	fmt.Printf("Blinds posted: SB %s pays %d, BB %s pays %d\n",
		g.Players[sbIndex].Name, sbAmount,
		g.Players[bbIndex].Name, bbAmount,
	)
}
