package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	foo()
	fmt.Println("Shall we do sth else?")
	baa()
	n, err := fmt.Println("and we exited!!")
	fmt.Println(n)  //print the number of character
	fmt.Println(err) //prinf the erorr message if any
}

func foo() {
	//lets print even numbers between 1 and 30
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func baa() {
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Print("*")
		}
		
		fmt.Print("\n")
	}
}

