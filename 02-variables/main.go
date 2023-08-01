package main

import "fmt"

// public variable
const Login string = "login"
const Number int = 200

func main() {
	var name string = "shaikr ahuman"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	var isLearnGoLang bool = true
	fmt.Println(isLearnGoLang)
	fmt.Printf("Variable is of type: %T \n", isLearnGoLang)

	var number int = 400
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

	// implicit type
	const coding = "GO"
	fmt.Println(coding)
	fmt.Printf("Variable is of type: %T \n", coding)

	// no var style
	website := "github.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// public variable print
	fmt.Println(Login)
	fmt.Printf("Variable is of type: %T \n", Login)

	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", Number)
}
