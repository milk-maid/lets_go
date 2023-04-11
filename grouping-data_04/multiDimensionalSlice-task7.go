package main

import (
	"fmt"
)

func main() {
	// Create a slice of a slice of string ([][]string).
	// Store the following data in the multi-dimensional slice:
	// "James", "Bond", "Shaken, not stirred"
	// "Miss", "Moneypenny", "Helloooooo, James."
	// Range over the records, then range over the data in each record

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xy := [][]string{x, y}

	// Todd Style
	for i, xy := range xy {
		fmt.Println("record: ", i)
		for j, val := range xy {
			fmt.Printf("\t at index: %v \t the value is: %v\n", j, val)
		}
		// fmt.Printf("at %d we have %v \n", i, v)
	}
	// Roqbell on this
	for i := 0; i < len(xy); i++ {
		fmt.Println("record: ", i)
		for j := 0; j < len(xy[i]); j++ {
			fmt.Printf("\t at index: %d \t the value is: %v\n", j, xy[i][j])
			// fmt.Printf("%v\n", xy[i][j])
		}
	}
	fmt.Println("Hello World!")
}

