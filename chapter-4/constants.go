package main

import "fmt"

func main() {
    // The const x's values cannot be changed later.
    const x string = "Hello World"
    fmt.Println(x)

    // Wont work:
    // x = "Hello Earth"
    // fmt.Println(x)
}
