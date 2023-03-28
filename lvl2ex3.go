package main

import "fmt"

// Create TYPED and UNTYPED constants. Print the values of the constants.
const (
	a = 12
	b = 13.56
	c = "Agboola Folawiyo"
)
const (
	m string  = "Bello Ologo"
	n int8    = 123
	o float64 = 234.768
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)

}
