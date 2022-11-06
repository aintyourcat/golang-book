/*
Cryptographic hash functions are similiar to their non-crytographic countrparts,
but they have the added property of being hard to reverse.
Given the cryptographic hash of a set of data, it's extremely difficult to
determine what made the hash.
These hashes are often used in security apps.
*/

package main

import (
	"crypto/sha1"
	"fmt"
)

/*
	This example is very similiar to the `crc32` one, because both `crc32` and
	`sha1` implement the hash.Hash interface.
	The main difference is that whereas `crc32` computes a 32 bit hash, `sha1`
	computes a 160 bit hash. There is no native type to represent a 160 bit number,
	so we use a slice of 20 bytes instead.
*/

func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
