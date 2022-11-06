package main

import (
	"fmt"
	"strings"
)

func main() {
	contains := strings.Contains("test", "e")
	fmt.Println("contains:", contains)

	count := strings.Count("test", "s")
	fmt.Println("count:", count)

	hasPrefix := strings.HasPrefix("test", "tes")
	fmt.Println("hasPrefix:", hasPrefix)

	hasSuffix := strings.HasSuffix("test", "est")
	fmt.Println("hasSuffix:", hasSuffix)

	index := strings.Index("test", "e")
	fmt.Println("index:", index)

	join := strings.Join([]string{"t", "e", "s", "t", "123"}, "-")
	fmt.Println("join:", join)

	repeat := strings.Repeat("i", 10)
	fmt.Println("repeat:", repeat)

	replace := strings.Replace("bee!", "e", "o", 2)
	fmt.Println("replace:", replace)

	split := strings.Split(join, "e")
	fmt.Println("split:", split)

	toLower := strings.ToLower("HELLO!")
	fmt.Println("toLower:", toLower)

	toUpper := strings.ToUpper(toLower)
	fmt.Println("toUpper:", toUpper)

	slice := []byte("hello")
	fmt.Println("slice:", slice)

	str := string([]byte{'h', 'e', 'l', 'l', 'o'})
	fmt.Println("str:", str)
}
