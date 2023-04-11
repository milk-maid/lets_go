package main

import (
	"fmt"
)

func main() {
	s := "Hello"
	fmt.Println(s)
	bs := []byte(s)
	fmt.Println("The slice bytes array of the above string is: ", bs)

	n := bs[0]
	fmt.Println("The value at index zero of the string in byte is: ", n)
	m := bs[1]
	fmt.Println("at index one we have:", m)
	fmt.Printf("the type of the value is : %T\n", m)
	fmt.Printf("the value in binary is : %b\n", m)
	fmt.Printf("the value in hex is : %#x\n", m)
}

