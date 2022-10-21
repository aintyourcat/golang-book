package main

import "fmt"

func swap(firstNumber *int, secondNumber *int) {
    *firstNumber, *secondNumber = *secondNumber, *firstNumber
}

func main() {
    firstNumber := 15
    secondNumber := 9

    fmt.Println("Before swap:", firstNumber, secondNumber)

    swap(&firstNumber, &secondNumber)

    fmt.Println("After swap:", firstNumber, secondNumber)
}
