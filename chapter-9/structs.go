package main

import (
	"fmt"
)

/*
A struct is a type which contains named fields.
The type keyword introduces a new type. It followed by the name of the type,
the keyword struct to indicate that we are defining a struct type
and a list of fields of curly braces.
Each field has a name and a type.
*/
// type Circle struct {
//     x float64
//     y float64
//     r float64
// }
// Or:
type Circle struct {
	x, y, r float64
}

func main() {
	/*
	   Creating an instance.
	   Like with other types, the following line create a local Circle variable
	   that is by default set to zero. For a struct zero means
	   each of the fields is set to their corresponding zero value
	   (0 for ints, 0.0 for floats, "" for strings, nil for pointers, ...).
	*/
	// var c Circle

	/*
	   The following allocates memory for all the fields,
	   sets each of them their zero value, and returns a pointer. (*Circle)
	*/
	// c := new(Circle)

	// Giving each fields a value:
	// c := Circle{x: 0, y: 0, r: 5}
	// If we know the order the fields were defined, we can omit the field names:
	c := Circle{0, 0, 5}

	fmt.Println(c)
}
