package main

import "fmt"

func main() {
	// if statement in action: create a program that uses “else if” and “else”.
	name := "Roqeeb"
	if name == "Bello" {
		fmt.Println("Correct it is ", name)
	} else if name == "Santorini" {
		fmt.Println("Yes! It is Santorini")
	} else {
		fmt.Println("Nope! actually the name is: ", name)
	}
}
