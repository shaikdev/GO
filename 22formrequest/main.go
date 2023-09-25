package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	const myUrl string = "https://lco.dev:3000"
	fmt.Println("Form request api...")
	data := url.Values{}
	data.Add("name", "shaik")
	data.Add("email", "shaikdev76@gmail.com")
	data.Add("age", "24")
	postData, err := http.PostForm(myUrl, data)
	isErr(err)
	defer postData.Body.Close()
	readData, readErr := ioutil.ReadAll(postData.Body)
	isErr(readErr)
	fmt.Println("The form data is: ", string(readData))

}

func isErr(err error) {
	if err != nil {
		panic(err)
	}
}
