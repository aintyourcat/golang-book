/*
Package declaration. Every Go program must start with it.
Packages are Go's way of organizing and reusing code.
There are two types of Go programs: executables and libraries.
Executables are the kinds of programs that we can run directly from the terminal.
Libraries are collections of code that packaged together, so they can be used in other programs.
*/
package main

/*
We use the import keyword to include code from other packages in our program.
The imported package's name, fmt, is sorrounded by double quotes.
The use of double quotes like that is known as a "string literal", which is a type of "expression".
*/
import "fmt"

/*
Comments are ignored by the Go compiler.
Programmer or whoever picks up the source code take advantage of it.
*/
// this is a comment

/*
Parameter and return type are optional.
The function main is the one that gets called when we execute the program.
*/
func main() {
    /*
    A statement.
    Invoking the Println function of the fmt package with a single argument.
    */
    // fmt.Println("Hello World")
    fmt.Println("Hello my name is behappier!")
}
