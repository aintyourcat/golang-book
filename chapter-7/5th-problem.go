package main

import "fmt"

func fibonacci(number uint) uint {
    if number == 0 {
        return 0
    } else if number == 1 {
        return 1
    } else {
        return fibonacci(number - 1) + fibonacci(number - 2)
    }
}

func main() {
    fmt.Println("fibonacci(0):", fibonacci(0))
    fmt.Println("fibonacci(1):", fibonacci(1))
    fmt.Println("fibonacci(2):", fibonacci(2))
    fmt.Println("fibonacci(3):", fibonacci(3))
}
