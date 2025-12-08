package main

import (
	"fmt"
	"poker-guru/engine"
	//"poker-guru/model"
)

func main() {
	fmt.Println("Welcome to Poker Guru!")

	deck := engine.NewDeck()
	for i := 0; i < 5; i++ {
		drawn1, error1 := engine.DrawCard(deck)
		if error1 != nil {
			panic(error1)
		}
		drawn2, error2 := engine.DrawCard(deck)
		if error2 != nil {
			panic(error2)
		}
		fmt.Println(drawn1, drawn2)
	}
}
