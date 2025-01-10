package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch class")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Random no. : ",diceNumber)

	switch diceNumber{
		case 1:
			fmt.Println("Dice value 1")
		case 2:
			fmt.Println("Dice value 2")
		case 3:
			fmt.Println("Dice value 3")
		case 4:
			fmt.Println("Dice value 4")
		case 5:
			fmt.Println("Dice value 5")
		case 6:
			fmt.Println("Dice value 6 and have 1 more chance")
		default:
			fmt.Println("What was that")
	}
}