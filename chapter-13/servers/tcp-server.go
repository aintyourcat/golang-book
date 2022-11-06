package main

/*
The `encoding/gob` package makes it easy to encode Go values
so that other Go programs (or the same Go program in this case)
can read them.

Additional encodings are available in packages underneath `encoding`
(like `encoding/json`) as well in 3rd party packages. (for example
we could use `labix.org/v2/mgo/bson` for bson support)
*/

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client(msg string) {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send the message
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main() {
	msgs := []string{
		"Hello World!",
		"Hello Mom!",
		"Hello Friends!",
	}

	go server()

	for _, msg := range msgs {
		go client(msg)
	}

	var input string
	fmt.Scanln(&input)
}
