/*
The preferred way to handle concurrency and synchronization
in Go is through goroutines and channels.

However Go does provide more traditional multithreading routines
in the `sync` and `sync/atomic` packages.

Great care should be taken when using mutexes or the synchronization
primitives provided in the `sync/atomic` package.

Traditional multithreaded programming is difficut; it's easy to
make mistakes and those mistakes are hard to find, since they may
depend on a very specific, relatively rare, and difficult to reproduce
set of circumstances.

One of Go's biggest strenghts is that the concurrency features it
provides are much easier to understand and use properly than
threads and locks.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		A mutex (mutual exclusion lock) locks a section of code to
		a single threaad at a time and is used to protect shared
		resources from non-atomic operations.
	*/
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			/*
				When the mutex (`m`) is locked any other attempt to lock it
				will block until it unlocked.
			*/
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
