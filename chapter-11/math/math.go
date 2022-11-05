/*
An important part of high quality software is code reuse --
embodied in the "Don't Repeat Yourself" principle.

Functions are the first layer we turn to allow code reuse.
Another mechanism Go provides for code reuse is packages.

Go's standard distribution provides a package named "math",
however since Go's package can be hierarchical we are safe to
use the same name for our package. (The real `math` package is just `math`,
ours is `golang-book/chapter-11/math`)

Package names match the folders they fall in.
(There's way around this, but it's a lot easier if we stay with this pattern)

When we import this package we use the full name, but inside this file
we only use the last part of the name. We also use the short name (`math`)
when we reference functions from our package.
*/
package math

/*
In Go if something starts with a capital letter that means
other package (and programs) are able to see it.
If we had named the following function `average` our `main` program
wouldn't have been able to see it.

It's a good practice to only expose the parts of our package that we want other
packages using and hide everything else.
*/

/*
Go has the ability to automatically generate documentation for packages we write
similiar way to the standard package documentation.

Run: `GO111MODULE=auto go doc golang-book/chapter-11/math Average` to see the documentation
of the following function. (It'll show the comment before the function)
*/

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the max value in a series of numbers
func Max(xs []float64) float64 {
	max := xs[0]

	for _, x := range xs {
		if x > max {
			max = x
		}
	}

	return max
}

// Finds the min value in a series of numbers
func Min(xs []float64) float64 {
	min := xs[0]

	for _, x := range xs {
		if x < min {
			min = x
		}
	}

	return min
}
