package main

import (
	"fmt"
)

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
  fmt.Println("the bytes array: ", bs)
  fmt.Printf("datatype of bytes in go: %T\n", bs)
  fmt.Printf("the byte in hex value: %#x\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")
	
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
