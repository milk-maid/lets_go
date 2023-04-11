package main

import (
	"fmt"
)

func main() {
	// ●Using a COMPOSITE LITERAL:
	// ● create an ARRAY which holds 5 VALUES of TYPE int
	// ● assign VALUES to each index position.
	// ● Range over the array and print the values out.
	// ● Using format printing
	// ○ print out the TYPE of the array
	var arr [5]int
	arr[0] = 12
	arr[1] = 23
	arr[2] = 34
	arr[3] = 45
	arr[4] = 56
	for i, v := range arr {
		fmt.Printf("at index %d there's value %d \n", i, v)
	}
	fmt.Printf("the type of the array is %T\n", arr)
	// array := [5]int{4,3,5,6,8}

	fmt.Println("Hello World")
}

