package main

import "fmt"

func main() {
    // A variable is a storage location, with a specific type and an associated name.

    /*
    Assign the string literal "Hello World" to a variable with name x with string type.
    var x string = "Hello World"
    Or:
    */
    var x string
    x = "Hello World"
    fmt.Println(x)

    // Variables can change their value throughout the lifetime of a program.
    x = "Hello Earth"
    fmt.Println(x)

    x = "Hello "
    x += "Moon" // Shorter concatenation statement
    fmt.Println(x)

    // String equality
    fmt.Println(`"hello" == "world" =`, "hello" == "world")
    fmt.Println(`"hello" == "hello" =`, "hello" == "hello")

    /*
    Shorter variable declaration & assigment statement.
    The variable's type is unspecified because it's unneccesary,
    the Go compiler will infer the type based on the literal value we assign to the variable.
    Generally, we should use this shorter form whenever possible.
    */
    y := "Hi World"
    fmt.Println(y)

    // Type inference is also working while using the var keyword:
    var z = 7
    fmt.Println(z)

    // Pick names which clearly describe the variable's purpose.
    // a := "Max" // Bad
    // name := "Max" // Better
    dogsName := "Max" // Best
    fmt.Println("My dog's name is", dogsName)
}
