package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is switch case statement")
	// rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	// switch case
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
		// fallthrough means the case has 2 run case 2 and case 3
		// the case doesn't have fallthrough the case 2 only run
		fallthrough
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that")
	}
}
