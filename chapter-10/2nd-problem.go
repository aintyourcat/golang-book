package main

import "time"

func Sleep(duration time.Duration) {
	c := time.After(duration)

	<-c
}

func main() {
	Sleep(time.Second)
}
