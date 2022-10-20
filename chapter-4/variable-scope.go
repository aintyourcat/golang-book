package main

import "fmt"

// The main and f functions have access to this variable.
y := "Hello Earth"

func main() {
    /*
    Only the main function has access to this variable.
    The range of places where we are allowed to use a variable is called the scope of a variable.
    Go is lexically scoped using blocks,
    it means that a variable exists within the nearest curly braces (a block),
    including any nested curly braces,
    but not outside them.
    */
    x := "Hello World"
    fmt.Println(x)
}

func f() {
    // The Go compiler will yell that the x variable is undefined.
    fmt.Println(x)
    fmt.Println(y)
}
