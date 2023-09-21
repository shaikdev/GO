package main

import "fmt"

func main() {
	fmt.Println("if else statements")
	var result string
	const number int = 3
	if number < 5 {
		result = "The number is less than 5"
	} else if number > 10 {
		result = "The number is greater than 10"
	} else {
		result = "The number is less than 10"
	}
	fmt.Println("The result is: ", result)

	// another method
	var numResult string
	if num := 6; num < 5 {
		numResult = "The num is less than 5"
	} else {
		numResult = "The num is greater than 10"
	}

	fmt.Println("The num is :", numResult)

	if 8%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The num is odd")
	}

}
