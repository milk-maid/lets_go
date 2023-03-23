package main

import (
	"fmt"
)

func main() {
  fmt.Println("Hello, playground")
	foo()
	fmt.Println("Shall we do sth else?")
  baa()
  fmt.Println("and we exited!!")
}


func foo() {
	//lets print even numbers between 1 and 19
	for i := 0; i < 19; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func baa() {
  for i := 0; i <= 5; i++ {
    for j := 0; j <= 5; j++ {
      fmt.Println("*")
    }
  }
}
