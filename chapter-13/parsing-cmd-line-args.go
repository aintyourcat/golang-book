package main

import (
	"flag"
	"fmt"
	"math/rand"
)

/*
When we invoke a command it's possible to pass the commands arguments.
`go run myfile.go`: run and myfile.go are arguments.
We can also pass flags to a command: `go run -v myfile.go`

The `flag` package allows us to parse arguments and flags sent to our program.
*/

// This program accept the -max flag.
func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
	/*
		Any additional non-flag args can be retrieved with
		flag.Args() which returns a []string
	*/
	fmt.Println(flag.Args())
}
