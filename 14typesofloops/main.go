package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// like forEach
	for index, value := range days {
		fmt.Printf("The index is %v and value is %v\n", index, value)
	}

	var number int = 1

	// like while loop
	for number < 10 {
		if number == 5 {
			// break stop and make out of loop
			// break
			number++
			// continue allows to resume loop
			continue
		}
		if number == 7 {
			goto jumping
		}

		fmt.Println("The values is: ", number)
		number++
	}

jumping:
	fmt.Println("Jumping...")

}
