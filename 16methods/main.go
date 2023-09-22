package main

import (
	"fmt"
)

type User struct {
	name  string
	email string
	age   int
}

func main() {
	user := User{"shaik", "shaikr1727@gmail.com", 24}
	fmt.Printf("My name is %v, My email id is %v, My age is %v\n", user.name, user.email, user.age)
	user.user()
	fmt.Println(user)
}

func (u User) user() {
	fmt.Println("This function for change a email id")
	u.email = "shaikdev76@gmail.com"
	fmt.Printf("My name is %v, My email id is %v, My age is %v\n", u.name, u.email, u.age)
}
