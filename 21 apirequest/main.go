package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	const url string = "https://lco.dev:3000/learn?coursename=reactjs"
	fmt.Println("api request")
	body := strings.NewReader(`
	{
		"name":"shaik"
		"age"24,

	}
	`)
	fmt.Println("Request body is: ", body)

	response, err := http.Post(url, "application/json", body)
	isError(err)
	fmt.Println("The post response is: ", response)
	defer response.Body.Close()
	// var convertString strings.Builder
	content, contentErr := ioutil.ReadAll(response.Body)
	isError(contentErr)
	// byteContent, _ := convertString.Write(content)
	// fmt.Println("The content is: ", convertString.String())
	fmt.Println("The content is: ", string(content))

}

func isError(err error) {
	if err != nil {
		panic(err)
	}
}
