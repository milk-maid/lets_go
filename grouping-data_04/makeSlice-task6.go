package main

import (
	"fmt"
)

// SLICE APPEND n DELETE

func main() {
	// Create a slice to store the names of all of the states in the United States of America. Use make
	// and append to do this. Goal: do not have the array that underlies the slice created more than
	// once. What is the length of your slice? What is the capacity? Print out all of the values, along
	// with their index position, without using the range clause. Here is a list of the 50 states:
	x := make([]string, 50, 50)
	usa := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`,
		` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`,
		` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`,
		` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`,
		` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i, v := range usa {
		x[i] = v
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	fmt.Println("Hello World!")
}

