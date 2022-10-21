package main

import "fmt"

/*
Pointers reference a location in memory where a value is stored
rather than the value itself. (They point to something else)
Pointers are rarely used with Go's builtin types,
but they are extremely useful when paired with structs.
*/

/*
By using a pointer (*int) the zero function is able to modify the original variable.
In Go a pointer is represented using the * character followed by the type of the stored value.
xPtr is a pointer to an int.
*/
func zero(xPtr *int) {
    /*
    * is also used to "dereference" pointer variables.
    Deferencing a pointer gives us access to the value the pointers point to.
    By *xPtr = 0 we are saying "store 0 in the memory location xPtr refers to"
    */
    fmt.Println("*xPtr:", *xPtr) // The stored value.
    *xPtr = 0
    
    /*
    Instead, if we try xPtr = 0 we will get a compiler error because xPtr is not an int,
    it's a *int, which can only be given another *int.
    */
    fmt.Println("xPtr:", xPtr) // The location in memory where the value stored.
    // xPtr = 0
}

func main() {
    x := 5
    fmt.Println("x:", x)
    
    /*
    We use the & operator to find the address of a variable.
    &x returns a *int (pointer to an int) because x is an int.
    This is what allow us to modify the original variable.
    &x in main and xPtr in zero refer to the same memory location.
    */
    fmt.Println("&x:", &x)
    zero(&x)

    fmt.Println("x:", x)
}
