package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {

	user := User{"shaik", "shaikrdev76@gmail.com", 24}

	// we make key and value pairs
	userKey := reflect.TypeOf(user)
	userValue := reflect.ValueOf(user)

	keyAndValue := make(map[string]interface{})

	// TODO: set key and value to keyAndValue
	for i := 0; i < userKey.NumField(); i++ {
		fieldKey := userKey.Field(i).Name
		fieldValue := userValue.Field(i).Interface()
		keyAndValue[fieldKey] = fieldValue
	}

	fmt.Println("The key and pair value is: ", keyAndValue)

	for key, value := range keyAndValue {
		fmt.Printf("The key %v and value is %v", key, value)
	}

}
