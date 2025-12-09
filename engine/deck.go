package engine

import (
	"fmt"
	"math/rand"
	"poker-guru/model"
	"time"
)

type Deck struct {
	cards []model.Card
}

func NewDeck() *Deck {
	cards := make([]model.Card, 0, 52)

	for s := model.Club; s <= model.Spade; s++ {
		for r := model.RankTwo; r <= model.RankAce; r++ {
			card := model.Card{
				Suit: s,
				Rank: r,
			}
			cards = append(cards, card)
		}
	}

	//shuffle
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return &Deck{
		cards: cards,
	}
}

func (d *Deck) DrawCard() (model.Card, error) {
	if len(d.cards) == 0 {
		if len(d.cards) == 0 {
			return model.Card{}, fmt.Errorf("deck is empty")
		}
	}

	// take the last card in the slice
	idx := len(d.cards) - 1
	card := d.cards[idx]

	// shrink the slice by one
	d.cards = d.cards[:idx]

	return card, nil
}
