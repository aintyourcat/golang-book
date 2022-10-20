package main

import "fmt"

func main() {
    var (
        x string = "Hello World"
        y int = 10
        z bool = true
    )

    const (
        a string = "Hello Earth"
        b int = 7
        c bool = false
    )

    fmt.Println(x, y, z, a, b, c)
}
