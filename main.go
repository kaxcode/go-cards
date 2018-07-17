package main

import (
	"fmt"
)

func main() {
	// only use := to define new variables
	card := "Aces of Spades"
	// reassign value of variable using = not :=
	card = "Five of Diamonds"
	fmt.Println(card)
}
