package main

import (
	"fmt"
	"math"
)

/*
We were able to name the Rectangle's area method
the same thing as the Circle's area method.
This was no accident. In both real life and in programming,
relationships like these are commonplace.
Go has a way of making these accidental similiarities explicit
through a type known as an Interface.
*/

/*
Like a struct, an interface is created using the type keyword,
followed by a name and the keyword interface.
But instead of defining fields, we define a "method set".
A method set is a list of methods that a type must have in order to
implement the interface.
In our case both Rectangle and Circle have area methods which return float64s
so both types implement the Shape interface.
*/
type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

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

// Using interface as an argument to a function
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// Using interface a field
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (c *Circle) perimeter() float64 {
	panic("Not Implemented")
}

func (r *Rectangle) perimeter() float64 {
	panic("Not Implemented")
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	// Unsure why i have to use &c and &r instead of c and r.
	m := MultiShape{[]Shape{&c, &r}}

	fmt.Println(totalArea(&c, &r))
	fmt.Println(m.area())

	c.perimeter()
	r.perimeter()
}
