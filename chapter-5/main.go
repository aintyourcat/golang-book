package main

import "fmt"

func main() {
    // for allows us to repeat a list of statements (a block) multiple times.
    // i := 1
    // for i <= 10 {
    //     fmt.Println(i)
    //     i += 1
    // }
    // Or:
    for i := 1; i <= 10; i++ {
        /*
        if allows us to run different blocks based on a condition.
        The first condition that evaluates to true will have it's associated block executed.
        */
        if i % 2 == 0 {
            fmt.Println(i, "even")
        } else {
            fmt.Println(i, "odd")
        }
    }

    j := 3

    /*
    switch similiar to if,
    however it compares an expression's value with one or more expressions values.
    The first matched expression's value will have it's associated statement(s) executed.
    */
    switch j {
        case 0: fmt.Println("Zero")
        case 1: fmt.Println("One")
        case 2: fmt.Println("Two")
        case 3: fmt.Println("Three")
        case 4: fmt.Println("Four")
        case 5: fmt.Println("Five")
        default: fmt.Println("Unknown Number")
    }
}
