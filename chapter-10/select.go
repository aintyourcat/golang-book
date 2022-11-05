package main

/*
`select` works like `switch` but for channels. It picks the first channel that is ready
and receive from it (or send to it). If more than one of the channels are ready then it
randomly picks which one to receive from. If none of the channels are ready, the statement
blocks until one become available.
*/

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Millisecond * 500):
				/*
					time.After creates a channel and after the given duration will
					send current time on it. (We didn't save it in a variable)
				*/
				fmt.Println("timeout")
			default:
				// The default case happens immediately if none of the channels are ready.
				fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
