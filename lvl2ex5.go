package main

import (
	"fmt"
)

// Create a variable of type string using a raw string literal. Print it.

func main() {
	boo := `Mudashiru 
  Lawal was a great 
  player
  He made history at a very tender age
  "What a great man he is!"`
	fmt.Println(boo)
	fmt.Printf("%s\n", boo)
	fmt.Println("Hello!")
}
