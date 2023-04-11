package main

import "fmt"

func main() {
  var m [58]string

  for i := 65; i <= 122; i++ {
    m[i - 65] = string(m)
  }
  fmt.Println(m)
  fmt.Println(m[41])
}
