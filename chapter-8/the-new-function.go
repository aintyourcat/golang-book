package main

import "fmt"

func one(xPtr *int) {
    *xPtr = 1
}

func main() {
    /*
    The built-in new function is another way to get a pointer.
    new takes a type as an argument,
    allocates enough memory to fit a value of that type,
    and returns a pointer to it.
    In some programming languages there is a significant difference between using new and &,
    with greate care being needed to eventually delete anyting created with new.
    Go isn't like this, it's a GARBAGE COLLECTION PROGRAMMING LANGUAGE
    which means memory is cleaned up automatically when nothing refers to it anymore.
    */
    xPtr := new(int)

    fmt.Println("xPtr:", xPtr)
    fmt.Println("*xPtr:", *xPtr)

    one(xPtr)

    fmt.Println("xPtr:", xPtr)
    fmt.Println("*xPtr:", *xPtr)
}
