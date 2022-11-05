package main

/*
Channels provide a way for two goroutines to communicate with one another
and synchronize their execution.
*/

import (
	"fmt"
	"time"
)

/*
When `pinger` attempts to send a message on the channel, it will wait until
`printer` is ready to receive the message. (this is known as blocking)
*/
func pinger(c chan<- string) { // c is a send-only channel, attempting to receive from `c` will result in a compiler error.
	for {
		c <- "ping" // means: send "ping" to channel `c`
	}
}

func ponger(c chan<- string) {
	for {
		c <- "pong" // means: send "pong" to channel `c`
	}
}

func printer(c <-chan string) { // c is a-receive only channel.
	for {
		msg := <-c // means: receive a message from channel `c` and store it in `msg`.
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	/*
		A channel is represented with the keyword `chan` followed by the type of things
		that are passed on the channel, string in our case.

		The `<-` operator is used to send and receive messages on the channel.

		We can specify a direction on a channel thus restricting it to either sending or receiving.

		A channel that doesnt have this restriction is known as bi-directional, `c` for example.

		A bi-directional channel can be passed to a function that takes send-only or receive-only
		channels, but the reverse is not true.
	*/
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
