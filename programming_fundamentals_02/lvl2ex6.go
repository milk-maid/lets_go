package main

import (
	"fmt"
)

// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

const (
	_ = iota
	y1 = 52 * iota
	y2 = 52 * iota
	y3 = 52 * iota
	y4 = 52 * iota
	)

func main() {
	fmt.Println("Number of weeks in a year is : ", y1)
	fmt.Println("Number of weeks in two years is : ", y2)
	fmt.Println("Number of weeks in three years is : ", y3)
	fmt.Println("Number of weeks in four years is : ", y4)
}
