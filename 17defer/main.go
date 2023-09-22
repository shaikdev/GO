package main

import "fmt"

func main() {
	// it mention defer the line will be moved to before closing bracket
	defer fmt.Println("WORLD")
	defer fmt.Println("ONE")
	defer fmt.Println("TWO")
	fmt.Println("HELLO")
	deferFunc()
	// o/p
	// HELLO, 43210, TWO, ONE, WORLD
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
