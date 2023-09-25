package main

import (
	"encoding/json"
	"fmt"
)

type Courses struct {
	CourseName string   `json:"course_name"`
	Price      int      `json:"price"`
	Website    string   `json:"website"`
	Password   string   `json:"-"`
	Tags       []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("create json type data")
	encodeJson()

}

func encodeJson() {
	data := []Courses{
		{"React", 200, "freecodecamp.com", "abc123", []string{"React.js", "JS"}},
		{"Angular", 200, "freecodecamp.com", "efg123", []string{"Angular.js", "Js"}},
		{"java", 200, "freecodecamp.com", "hij123", nil},
	}

	// convert to Json format
	jsonFormat, err := json.MarshalIndent(data, "", "\t")
	isError((err))
	fmt.Println("The json format data: ", string(jsonFormat))
}

func isError(err error) {
	if err != nil {
		panic(err)
	}
}
