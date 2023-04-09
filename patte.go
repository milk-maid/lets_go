package main

import "fmt"

func main() {
  i := 0 for i <= 7 {
    j := 0 for j <= 5 {
      fmt.Print("*") j++
    } fmt.Println() i++
  } fmt.Println("A job well done!")
}
