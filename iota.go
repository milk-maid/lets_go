package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
	d
	e
	f
)

const (
	m = iota
	n
	o
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
}
