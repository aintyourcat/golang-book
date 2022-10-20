package main

import "fmt"

/*
A function is an independent section of code that
maps zero or more input parameters to zero or more output parameters.
The average function take in a slice of float64s and return one float64.
Collectively the parameters and return type are known as the function's signature.
*/
func average(xs []float64) float64 {
    // panic("Not Implemented") // Causes a run time error.
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

// We can also name the return type
func f2() (r int) {
    r = 1
    return // r is automatically returned.
}

/*
Returning multiple values.
Multiple values are often used to return an error value along with the result.
*/
func f() (int, int) {
    return 5, 6
}

/*
Variadic function.
By using ... before the type name of the last parameter,
we indicate that it takes zero or more of those parameters.
In this case we take zero or more ints.
*/
func add (args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

// A function which returns another function.
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

// Recursion: a function which calls itself.
func factorial (x uint) uint {
    if x == 0 {
        return 1
    }

    return x * factorial(x-1)
}

func main() {
    xs := []float64{98, 93, 77, 82, 83}
    fmt.Println("average(xs):", average(xs))

    fmt.Println("f2():", f2())

    x, y := f()
    fmt.Println("x:", x)
    fmt.Println("y:", y)

    // We invoke the function with many ints as we want!
    fmt.Println("add(1, 2, 3):", add(1, 2, 3))
    // Or
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("add(numbers...):", add(numbers...))

    /*
    Closure: functions inside of functions.
    add is a variable with type func(int, int) int,
    a function that takes two ints and returns an int.
    */
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println("add(1, 1):", add(1, 1))

    nextEven := makeEvenGenerator()
    
    fmt.Println("nextEven():", nextEven()) // 0
    fmt.Println("nextEven():", nextEven()) // 2
    fmt.Println("nextEven():", nextEven()) // 4

    fmt.Println("factorial(0):", factorial(0))
    fmt.Println("factorial(1):", factorial(1))
    fmt.Println("factorial(2):", factorial(2))
}
