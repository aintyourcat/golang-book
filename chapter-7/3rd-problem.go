package main

import "fmt"

func findTheGreatestNumber(numbers ...float64) (greatestNumber float64) {
    greatestNumber = numbers[0]

    for _, value := range numbers {
        if value > greatestNumber {
            greatestNumber = value
        }
    }

    return
}

func main() {
    fmt.Println(findTheGreatestNumber(1, 2, 3, 4))
    fmt.Println(findTheGreatestNumber(90, 80, 70, 100, 60))
    fmt.Println(findTheGreatestNumber([]float64{-5, -6, -7, -8, -9}...))
}
