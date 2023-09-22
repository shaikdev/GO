package main

import (
	"fmt"
	"net/url"
)

var myUrl string = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {
	fmt.Println("This is handling url statements")
	// handling url
	urlParse, err := url.Parse(myUrl)
	checkErrorNil(err)
	// fmt.Println("schema", urlParse.Scheme)
	// fmt.Println("path", urlParse.Path)
	// fmt.Println("port", urlParse.Port())
	// fmt.Println("rawSchema", urlParse.RawQuery)

	query := urlParse.Query()
	fmt.Println("The query is: ", query)
	fmt.Println("The query value: ", query["coursename"])
	// print multiple query values
	for _, value := range query {
		fmt.Println("The query value is: ", value)
	}

	// another way to to make an url
	makUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	fmt.Println(makUrl)

}

func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
