package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input) // Scanf fills input with the number we enter.

    output := input * 2

    fmt.Println(output)
}
