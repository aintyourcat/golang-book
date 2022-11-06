package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

/*
	To define our own sort we create a new type (`ByName`) and
	make it equivalent to a slice of what we want to sort.

	The `Sort` function of the `sort` package takes a sort.Interface and sorts it.

	We then implement the 3 methods of that interface to our sort.
*/

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	/*
		Sorting our list then as easy as casting the list
		into our new type (the sort).
	*/
	sort.Sort(ByName(kids))
	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
