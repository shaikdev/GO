package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	const content string = "Hello this is the testing for golang file creation"
	file, err := os.Create("./mygolangfile.text")
	checkErrorNil(err)
	_, writeErr := io.WriteString(file, content)
	checkErrorNil(writeErr)
	file.Close()
	readFile("./mygolangfile.text")

}

func readFile(path string) {
	readFile, err := ioutil.ReadFile(path)
	checkErrorNil(err)
	fmt.Println("The file content is: ", string(readFile))
}

func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
