package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	for i, v := range x {
		fmt.Println(i, v)
	}
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
  x = append(x, 66, 77, 99, 6712, 01,123)
  y := []int{13, 14, 15, 17, 19, 20, 23}
  x = append(x, y...)
  fmt.Println(x)
  fmt.Println(x[2:5]) //slicing a slice
  x = append(x[:3], x[7:]...) //delete index 3 to 6
  fmt.Println(x) 

	fmt.Print("Thanks!\n")
}

