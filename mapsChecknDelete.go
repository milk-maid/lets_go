// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import "fmt"

func main() {
	myMap := map[string]int{"apple": 1, "banana": 2, "orange": 3}

	value, exists := myMap["apple"]
	if exists {
		fmt.Printf("The value of 'apple' is %d\n", value)
	} else {
		fmt.Println("'apple' does not exist in the map")
	}

  fmt.Println(myMap)

	// USE THIS DELETE STYLE
	if value, exists := myMap["apple"]; exists {
		fmt.Println("value: ", value)
		delete(myMap, "apple")
	}

	fmt.Println(myMap)

	fmt.Println("hello world")
}

