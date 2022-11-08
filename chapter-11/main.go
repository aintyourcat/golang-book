// Executable commands must always use `package main`
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
It seems that the book is still using an older method. (pre-modules)
Yes. This method of working with package relies on `GOPATH`.
The `GOPATH` env specifies the location of our Go workspace, which
defaults to $HOME/go in Unix. That's why we have to put this repo
(`golang-book`) under GOPATH/src.

A workspace is a directory with two directories at its root:
1. src contains Go source files, and
2. bin contains executable commands.

The src sub-directory typically contains multiple version control
repositories.

The bin sub-directory contains binaries built and installed by
the `go install`.

Obviously, we can use a different directory as our workspace, simply
by specifying it in `GOPATH` env.
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
