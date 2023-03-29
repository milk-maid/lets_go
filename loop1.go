package main

import "fmt"

var x int

func main() {
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Hello!!!!")

	y := 2
	for y < 13 {
		if y > 8 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("Franc!")

	z := 2
	for z < 30 {
		z++
		if z > 30 {
			break
		}
		if z%2 != 0 {
			continue
		}
		fmt.Println(z)
	}
	fmt.Println("Thanks!!!!")
}
