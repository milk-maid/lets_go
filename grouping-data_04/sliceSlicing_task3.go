package main

import (
	"fmt"
)

// SLICING A SLICE

func main() {
	// Using the code from the previous example, use SLICING
	// to create the following new slices which are then printed:
	// ● [42 43 44 45 46]
	// ● [47 48 49 50 51]
	// ● [44 45 46 47 48]
	// ● [43 44 45 46 47]

	sli := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(sli[:5])
	fmt.Println(sli[5:])
	fmt.Println(sli[2:7])
	fmt.Println(sli[1:6])

	fmt.Println("Hello World")
}

