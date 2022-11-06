package main

import (
	"container/list"
	"fmt"
)

/*
A doubly-linked list is a type of data structure that made of nodes,
each node of the list contains a value (1, 2, or 3 for example), and
pointers to the next and previous nodes.
*/

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
