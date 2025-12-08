package engine

import (
	"poker-guru/model",
	"fmt"
)

// Game represents the state of a poker game.
type Game struct {
	Deck	[]model.Card
	Players	[]model.Player
	Pot		int
}

// NewGame initializes a new game with players.
func NewGame(players []model.Player) *Game {
	return &Game{
        Deck:    initializeDeck(),
        Players: players,
        Pot:     0,
    }
}

//Start begins the poker game.
func (g *Game) Start() {
	fmt.Println("Starting the game...")
    shuffleDeck(g.Deck)
    dealCards(g.Players, g.Deck)

    // Implement game loop logic here
}

// initializeDeck creates a standard deck of 52 cards.
func initializeDeck() []model.Card {
    suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
    ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
    var deck []model.Card
    for _, suit := range suits {
        for _, rank := range ranks {
            deck = append(deck, model.Card{Suit: suit, Rank: rank})
        }
    }
    return deck
}

// shuffleDeck shuffles the deck of cards.
func shuffleDeck(deck []model.Card) {
    // Implement shuffling logic here
}

// dealCards deals cards to each player.
func dealCards(players []model.Player, deck []model.Card) {
    // Implement dealing logic here
}
