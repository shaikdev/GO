package main

import "fmt"

func main() {

	fmt.Println("This is go lang function statements")
	greeting()
	addNumber := addNumbers(5, 5)
	fmt.Println("The add numbers function value is: ", addNumber)
	proAdderValue, proAdderText := proAdder(2, 4, 6, 8, 2)
	fmt.Println("The proAdder functions value is: ", proAdderValue)
	fmt.Println("The proAdder function text is: ", proAdderText)

}

func greeting() {
	fmt.Println("This is greeting function")
}

func addNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(numbers ...int) (int, string) {
	total := 0
	for _, value := range numbers {
		total += value
	}
	// we passing two return statements like this
	return total, "The proAdder function completed"

}
