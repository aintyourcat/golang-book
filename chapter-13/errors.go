package main

import (
	"errors"
	"fmt"
)

func main() {
	// Creating our own errors
	err := errors.New("error message")

	fmt.Println(err)
}
