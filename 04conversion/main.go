package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	numberConversion, conversionErr := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if conversionErr != nil {
		fmt.Println(conversionErr)
	} else {
		fmt.Println("Thanks for rating", numberConversion)
	}
}
