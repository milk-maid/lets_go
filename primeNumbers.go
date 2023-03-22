package main

import "fmt"

func main() {
    var upperLimit int
    fmt.Print("Enter upper limit: ")
    fmt.Scanln(&upperLimit)

    for i := 2; i <= upperLimit; i++ {
        if isPrime(i) {
            fmt.Println(i)
        }
    }
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
