package main

import (
	"fmt"
)

// assigns an int to a variable
// ● prints that int in decimal, binary, and hex
// ● shifts the bits of that int over 1 position to the left, and assigns that to a variable
// ● prints that variable in decimal, binary, and hex

func main() {
	x := 42
	fmt.Printf("The value of x --\n in decimal: %d\n in binary %b\n in hex: %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("The bitwise shifted value of x --\n in decimal: %d\n in binary %b\n in hex: %#x\n", y, y, y)
	fmt.Println("Hello!")
}
