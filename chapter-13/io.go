package main

/*
The io package contains two main interface: `Reader` and `Writer`.
`Readers` support reading via the `Read` method.
`Writers` support writing via the `Write` method.
These interfaces are used in other packages.
*/

import (
	"bytes"
	"fmt"
)

func main() {
	/*
		We can use the `Buffer` struct from the `bytes` package to
		read or write to a []byte or string.
		A `Buffer` doesn't have to be initialized (The documentation say so) and
		supports both the `Reader` and `Writer` interfaces.
	*/
	var buf bytes.Buffer
	str := "test"

	// Writing:
	p := []byte(str)
	buf.Write(p)

	// Reading:
	// n := buf.Len()
	// p = make([]byte, n)
	// buf.Read(p)

	// fmt.Println(string(p))
	// Or, we can convert it into a []byte using buf.Bytes():
	fmt.Println(string(buf.Bytes()))
}
