package main

import "fmt"

func main() {
	// make(map[key]value)
	language := make(map[string]string)
	language["js"] = "javascript"
	language["py"] = "python"
	language["c"] = "c"
	language["rb"] = "ruby"

	fmt.Println("The language of list", language)
	fmt.Println("Js", language["js"])

	// loop
	for key, value := range language {
		fmt.Printf("The key %v, The value %v\n", key, value)
	}

}
