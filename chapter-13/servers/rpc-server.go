package main

/*
The `net/rpc` (remote procedure call) and `net/rpc/jsonrpc`
packages provide an easy way to expose methods so they can
be invoked over a network.
(rather than just in the program running them)
*/

import (
	"fmt"
	"net"
	"net/rpc"
)

/*
This program similiar to the TCP server one, except now
we created an object to hold all the methods we want to expose
and we call the `Negate` and `Add` methods from the client.
*/

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func (this *Server) Add(nums [2]int, reply *int) error {
	*reply = nums[0] + nums[1]
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var res1 int64
	err = c.Call("Server.Negate", int64(999), &res1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", res1)
	}

	var res2 int
	err = c.Call("Server.Add", [2]int{1, 2}, &res2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Add([1 2]) =", res2)
	}
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
