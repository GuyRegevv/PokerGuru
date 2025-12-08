package model

// Card represents a playing card with a suit and rank.
type Card struct {
    Suit string // "Hearts", "Diamonds", "Clubs", "Spades"
    Rank string // "2", "3", ..., "10", "J", "Q", "K", "A"
}

// Action represents a player's action in the game.
type Action struct {
    PlayerName string // Name of the player taking the action
    Move       string // "FOLD", "CALL", "RAISE", etc.
    Amount     int    // Amount involved in the action, if applicable
}

// Player represents a player in the game.
type Player struct {
    Name  string // Player's name
    Stack int    // Player's current stack size
    Hand  []Card // Player's current hand
}
