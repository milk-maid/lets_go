package main

import (
	"fmt"
)

func main() {
	// ●Using a COMPOSITE LITERAL:
	// ● create a SLICE of TYPE int
	// ● assign 10 VALUES
	// ● Range over the slice and print the values out.
	// ● Using format printing
	// ○ print out the TYPE of the slice

	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range slice {
		fmt.Printf("at index %d there's value %d \n", i, v)
	}
	fmt.Printf("the type of the slice is %T\n", slice)

	fmt.Println("Hello World")
}

