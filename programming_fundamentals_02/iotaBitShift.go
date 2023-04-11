package main

import (
	"fmt"
)

const ( 
	_ = iota
	kbi = 1 << (iota * 10)
	mbi = 1 << (iota * 10)
	gbi = 1 << (iota * 10)
)

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("kb in decimal: %d\t\t\t in binary: %b\n", kb, kb)
	fmt.Printf("mb in decimal: %d\t\t\t in binary: %b\n", mb, mb)
	fmt.Printf("gb in decimal: %d\t\t in binary: %b\n", gb, gb)
	fmt.Println("")
	fmt.Println("using iota for the increament!")
	fmt.Printf("kb in decimal: %d\t\t\t in binary: %b\n", kbi, kbi)
	fmt.Printf("mb in decimal: %d\t\t\t in binary: %b\n", mbi, mbi)
	fmt.Printf("gb in decimal: %d\t\t in binary: %b\n", gbi, gbi)
}
