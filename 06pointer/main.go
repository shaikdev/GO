package main

import "fmt"

func main() {
	myNumber := 25
	fmt.Println("Actual myNumber value  is :", myNumber)
	// using & to reference a value
	ptr := &myNumber
	fmt.Println("Actual pointer memory value is :", ptr)
	fmt.Println("Actual pointer  value is :", *ptr)
	// To change ptr  value
	*ptr = *ptr * 2
	fmt.Println("The modify ptr value :", *ptr)
	fmt.Println("Current myNumber values is :", myNumber)

}
