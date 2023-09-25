package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	coursename string
	price      int
	website    string
	tags       []string
}

func main() {
	fmt.Println("Consuming json...")
	decodeJson()

}

func decodeJson() {
	dataFromWeb := []byte(`
	{
		"course_name": "React",
		"price": 200,
		"website": "freecodecamp.com",
		"tags": [
				"React.js",
				"JS"]
	}
	`)

	var jsonValidIterface Course

	// check valid or not
	checkJsonValid := json.Valid(dataFromWeb)
	if checkJsonValid {
		fmt.Println(checkJsonValid)
		json.Unmarshal(dataFromWeb, &jsonValidIterface)
		fmt.Printf("%#v\n", jsonValidIterface)
	} else {
		fmt.Println("The json is invalid")
	}

	// like key pair value
	var keyPairValue map[string]interface{}
	json.Unmarshal(dataFromWeb, &keyPairValue)
	fmt.Printf("key pair value %#v\n", keyPairValue)

	// print the value
	for key, value := range keyPairValue {
		fmt.Printf("The key is %v and value is %v and type is %v\n", key, value, value)
	}

}
