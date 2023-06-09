package main

import (
	"fmt"
)

// SLICE APPEND n DELETE

func main() {
	// ● start with this slice
	// ○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// ● use APPEND & SLICING to get these values here which you should ASSIGN to a
	// variable “y” and then print:
	// ○ [42, 43, 44, 48, 49, 50, 51]
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)

	fmt.Println("Hello World!")
}

