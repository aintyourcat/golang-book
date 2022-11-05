package main

/*
Packages only make sense in the context of a separate program which uses them.
Without this separate program we have no way of using our package we create.

This program is using the `math` package we create.

If we wanted to use the real `math` package, Go allows us to use an alias
for our package, e.g: `import m "golang-book/chapter-11/math"` to
use m as the package's alias.
*/
import (
	"fmt"
	"golang-book/chapter-11/math"
)

/*
We have to run this as: `GO111MODULE=auto go run main.go`.
It seems that the book is still using an older method.
*/
func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	min := math.Min(xs)
	max := math.Max(xs)
	fmt.Println(avg)
	fmt.Println(min)
	fmt.Println(max)
}
