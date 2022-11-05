package main

/*
Concurreny is making progress on more than one task simultaneously.
According to Rob Pike, concurrency is not parallelism:
concurrency is about dealing with lots of things at once but
 parallelism is about doing lots of things at once.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		/*
			Goroutine is a function that is capable of running concurrently with other functions.
			To create a goroutine we use the keyword `go` followed by function invocation.
			This program consists of several goroutines: the first one is implicit and is the main function itself,
			the other goroutines are created when we call `go f(i)`.
			Goroutines are lightweight and we can easily make thousands of them.
		*/
		go f(i)
	}
	/*
		Normally when we invoke a function our program will execute all statements in a function and
		then return to the next line following the invocations.
		With a goroutine we return immediately to the next line and don't wait for the function to
		complete. This is why the call to the `Scanln` function has been included, without it
		our program would exit before being given the opportunity to print all numbers.
	*/
	var input string
	fmt.Scanln(&input)
	fmt.Println("You typed: ", input)
}
