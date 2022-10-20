package main

import "fmt"

func main() {
    /*
    A call to the panic function cause a run time error.
    We can handle a run-time panic with the built-in recover function.
    recover stops the panic and returns the value that was passed to the call to panic.
    */

    // The following doesn't work because the call to panic immediately stops the execution of the function.
    // panic("PANIC!")
    // str := recover()
    // fmt.Println(str)

    // We have to pair it with defer instead:
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC!")

    /*
    It's called "panic" because it indicates a programmer error (e.g.
    attempting to access an index of an array that's out of bounds,
    forgetting to initialize a map, etc.) or
    an exceptional condition that there's no easy way to recover from.
    */
}
