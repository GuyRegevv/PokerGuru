package main

import (
	"fmt"
	"poker-guru/engine"
	"poker-guru/model"
)

func main() {
	fmt.Println("Welcome to Poker Guru!")

	players := []model.Player{
		{Name: "Alice", Stack: 100},
		{Name: "Bob", Stack: 100},
		{Name: "You", Stack: 100},
	}

	game := engine.NewGame(players)

	if err := game.StartHand(); err != nil {
		panic(err)
	}

	if err := game.DealHoleCards(); err != nil {
		panic(err)
	}
	if err := game.DealFlop(); err != nil {
		panic(err)
	}
	if err := game.DealTurn(); err != nil {
		panic(err)
	}
	if err := game.DealRiver(); err != nil {
		panic(err)
	}

	fmt.Println("Pot:", game.Pot)

	fmt.Println("Stacks after blinds:")
	for _, p := range game.Players {
		fmt.Println(p.Name, "stack:", p.Stack, "hand:", p.Hand)
	}

	fmt.Println("Board:", game.CommunityCards)
}
