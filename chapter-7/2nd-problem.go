package main

import "fmt"

func halve(number int) (int, bool) {
    return number / 2, number % 2 == 0
}

func main() {
    fmt.Println("halve(1):", halve(1))
    fmt.Println("halve(2):", halve(2))
    fmt.Println("halve(3):", halve(3))
}
