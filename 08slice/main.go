package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruits = []string{"Apple", "Orange"}
	fruits = append(fruits, "Banana", "Lemon")
	fmt.Println("Fruits array list: ", fruits)

	// slices
	sliceArray := append(fruits[1:])
	// its return except apple
	fmt.Println("slices array list: ", sliceArray)

	veg := []string{"Potato", "Tomato", "Garlic"}
	fmt.Println("This is veg array list: ", veg)
	// slice range between 0 and 2
	sliceVegArray := append(veg[0:2])
	fmt.Println("sliceVeg array list: ", sliceVegArray)

	// another way of approach
	arrayList := make([]int, 4)
	arrayList[0] = 923
	arrayList[1] = 456
	arrayList[2] = 245
	arrayList[3] = 779
	fmt.Println("arrayList: ", arrayList)

	arrayList = append(arrayList, 678)
	fmt.Println("append array list: ", arrayList)
	// is sorter or not return boolean
	fmt.Println("array list are sorted: ", sort.IntsAreSorted(arrayList))
	sort.Ints(arrayList)
	fmt.Println("array list are sorted: ", sort.IntsAreSorted(arrayList))
	fmt.Println("sort array list: ", arrayList)

}
