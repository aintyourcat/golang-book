package main

import "fmt"

func first() {
    fmt.Println("1st")
}

func second() {
    fmt.Println("2st")
}

func main() {
    /*
    defer schedules a function call to be run after the function completes.
    Basically, defer moves the call to second to the end of the function:
    */
    defer second()
    first()
    // Same as:
    // first()
    // second()

    /*
    defer is often used when resources need to be freed in some way.
    For example, when we open a file we need to make sure to close it later.
    f, _ := os.Open(filename)
    defer f.Close()
    This has 3 advantages:
    1. It keeps our Close call near our Open call so it's easier to understand.
    2. If our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before both of them.
    3. Deferred functions are run even if a run time panic occurs.
    */

    third()
}

func third() {
    defer fmt.Println("1st deferred call")
    fmt.Println("1st non-deffered call")

    defer fmt.Println("2nd deferred call")
    fmt.Println("2st non-deffered call")

    defer fmt.Println("3nd deferred call")
    fmt.Println("3st non-deffered call")
}
