package main;

import "fmt";

var y = 44;

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y);
	fmt.Printf("%b\n", y);
	fmt.Printf("%x\n", y);
	fmt.Printf("%#x\n", y);
}