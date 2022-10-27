package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

/*
area is a special type of function known as a method.
In between the func keyword and the name of the function (area),
we've added a "receiver".
The receiver is like a parameter - it has a name and a type -
but by creating the function in this way it allows us to
call the function using the . operator.
Since this function can only be used with the Circles we
can rename it to just area.
*/
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	/*
	   Calling the area method.
	   We no longer need the & operator in front of the c variable
	   (Go automatically knows to pass a pointer to the Circle for this method)
	*/
	fmt.Println(c.area())

	fmt.Println(r.area())
}
