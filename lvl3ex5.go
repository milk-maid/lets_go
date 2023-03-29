package main

import "fmt"

func main() {
	// Print out the remainder (modulus) which is found for each
	// number between 10 and 100 when it is divided by 4.

	for x := 10; x <= 100; x++ {
		mod := x % 4
		fmt.Println("the remainder of ", x, " divided by 4: ", mod)
	}
}

