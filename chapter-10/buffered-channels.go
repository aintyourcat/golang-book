package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Passing a second argument to the make function when creating a channel creates
		a buffered channel.

		The following creates a buffered channel with capacity of 1.
	*/
	c1 := make(chan int, 1)
	c2 := make(chan int)

	go func(c chan int) {
		fmt.Println("Does the following block?")
		c <- 2 // Blocks for about 5 seconds until the other side on the main goroutine is ready.
		fmt.Println("Yes it is.")
	}(c2)

	/*
		Normally channels are synchronous, both side of the channel will wait until the other side is ready.

		The buffered channel is asynchornous, sending or receiving a message will not wait unless the channel
		is already full.
	*/
	c1 <- 1 // Won't blocks
	fmt.Println(<-c1)

	time.Sleep(time.Second * 5)
	fmt.Println(<-c2)
}
