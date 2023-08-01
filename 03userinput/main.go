package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the rating for our coding")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
}
