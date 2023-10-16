package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var group = sync.WaitGroup{}

// go run -race test.go
func main() {
	var count int64
	count = 0
	totalGoRoutines := 10

	group.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			// lock

			//yield
			runtime.Gosched()

			// unlock
			fmt.Println("Count\t", atomic.LoadInt64(&count))
			group.Done()
		}()
		fmt.Println("Goroutine: ", runtime.NumGoroutine())
	}

	group.Wait()
	fmt.Println("count: ", count)

}
