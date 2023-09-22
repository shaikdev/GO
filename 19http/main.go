package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Make https request...")
	response, err := http.Get("https://lco.dev")
	checkErrorNil(err)
	defer response.Body.Close()
	readResponse, err := ioutil.ReadAll(response.Body)
	checkErrorNil(err)
	content := string(readResponse)
	fmt.Println("The response is: ", content)

}

func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
