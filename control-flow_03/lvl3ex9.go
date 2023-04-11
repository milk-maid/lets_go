package main

import (
	"fmt"
)

func main() {

	// Create a program that uses a switch statement with the switch expression specified as a
	// variable of TYPE string with the IDENTIFIER “favSport”.

	favSport := "surfing"

	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!!!")
	case "swimming":
		fmt.Println("do you have a pool???")
	case "surfing":
		fmt.Println("let's go to HAWAII!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral!")
	}
	fmt.Println("Hello!!!")
}

