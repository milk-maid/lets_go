package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mb := []string{"Sodiq", "Muri", "Banana", "Hazelnut"}
	fmt.Println(mb)

	xb := [][]string{jb, mb} //multidimensional slice
	fmt.Println(xb)
}

