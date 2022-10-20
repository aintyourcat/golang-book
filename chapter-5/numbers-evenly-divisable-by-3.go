package main

import "fmt"

func main() {
    for i := 100; i >= 1; i-- {
        if (i % 3 == 0) {
            fmt.Println(i)
        }
    }
}
