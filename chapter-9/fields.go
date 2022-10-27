package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// func circleArea(c Circle) float64 {
//     return math.Pi * c.r*c.r
// }

/*
Arguments are always copied in Go.
If we want to modify one of the fields inside of the circleArea function, it would not modify the original variable.
Because of this we would typically write the function like this
*/
func circleArea(c *Circle) float64 {
	/*
	   https://go.dev/ref/spec#Selectors
	   c.r is a shorthand of (*c).r
	*/
	return math.Pi * c.r * c.r
}

func main() {
	var c Circle = Circle{0, 0, 5}

	// We can access the fields using the . operator:
	c.x = 10
	c.y = 5
	fmt.Println(c.x, c.y, c.r)

	// fmt.Println(circleArea(c))
	fmt.Println(circleArea(&c))
}
