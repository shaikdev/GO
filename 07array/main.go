package main

import (
	"fmt"
)

func main() {
	// we mention array has contain only 5 data
	var myArray [5]string
	myArray[0] = "one"
	myArray[1] = "two"
	myArray[2] = "three"
	myArray[3] = "four"
	myArray[4] = "five"
	fmt.Println("array list: ", myArray)
	fmt.Println("array ist length: ", len(myArray))

	//   another way & without array size
	var arrayWithoutSize = []string{"shaik", "rahuman"}
	fmt.Println("array lis without size: ", arrayWithoutSize)

}
