package main

import "fmt"

func main() {
	var myArray = []string{"Javascript", "Python", "Ruby", "Java", "C++"}
	// am removed java by index number
	const indexNumber int = 3
	myArray = append(myArray[:indexNumber], myArray[indexNumber+1:]...)
	fmt.Println("The sliced array of list", myArray)
}
