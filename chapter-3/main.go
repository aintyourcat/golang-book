package main

import "fmt"

func main() {
    fmt.Println("1 + 1 =", 1 + 1) // The numeric literal 1's data type is int.
    fmt.Println("1 + 1 =", 1.0 + 1.0) // .0 used to tell Go that this is a floating-point number.
    
    fmt.Println(len("Hello World"))
    /*
    Prints 101 instead of e, since the character is represented by a byte,
    a byte is an integer (byte aka int8).
    */
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")

    // A boolean value (named after George Boole) is a special 1 bit integer to represent true and false.
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
