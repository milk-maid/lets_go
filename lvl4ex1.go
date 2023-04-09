package main

import (
	"fmt"
)

func main() {
	// ● Using a COMPOSITE LITERAL:
	// ● create an ARRAY which holds 5 VALUES of TYPE int
	// ● assign VALUES to each index position.
	// ● Range over the array and print the values out.
	// ● Using format printing
	// ○ print out the TYPE of the array
	var x [5]int
	x[0] = 12
	x[1] = 54
	x[2] = 3
	x[3] = 42
	x[4] = 75

	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("The type of the array is::: %T\n\n", x)

	// /////////ANOTHER ANGLE/// SLICE /////

	h := [5]int{311, 4, 3, 6, 1}
	fmt.Println(h)
	fmt.Printf("type of h is %T\n\n", h)

	m := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(m)
	for i, v := range m {
		fmt.Println(i, v)
	}
	fmt.Printf("The type of this SLICE is::: %T\n", m)

}

