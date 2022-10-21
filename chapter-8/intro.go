package main

import "fmt"

/*
When we call a function that takes an argument,
that argument is copied to the function.
*/

func zero(x int) {
    x = 0
}

func main() {
    x := 5
    /*
    the zero function will not modify the original x variable.
    But what if we want to?
    One way to do this is to use a special data type known as a pointer.
    */
    zero(x)
    fmt.Println(x)
}
