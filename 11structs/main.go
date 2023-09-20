package main

import "fmt"

type User struct {
	name       string
	email      string
	age        int
	isVerified bool
}

func main() {
	userData := User{"Shaik", "shaikdev@gmail.com", 24, true}
	fmt.Println("This is an user data", userData)

	// we print like key and value
	fmt.Printf("My name is %v, My email is %v, My age is %v, My status is %v\n", userData.name, userData.email, userData.age, userData.isVerified)

	// we describe the data key and value pair
	fmt.Printf("My user data %+v\n", userData)

}
