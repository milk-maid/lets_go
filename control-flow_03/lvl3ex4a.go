package main

import "fmt"

func main() {
	// Create a for loop using this syntax
	// ● for condition { }
	// Have it print out the years( birthyear == btyr ) you have been alive.
	btyr := 1993
	for {
    if btyr > 2023 {
      break
    }
		fmt.Println(btyr)
    btyr++
	}
}
