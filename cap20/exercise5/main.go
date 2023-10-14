package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var group = sync.WaitGroup{}
var count int32

func main() {
	count = 0

	numberOfGoroutines := 10

	group.Add(numberOfGoroutines)

	fmt.Println("goroutines: ", runtime.NumGoroutine())

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {

			atomic.AddInt32(&count, 1)

			fmt.Println("Count\t", atomic.LoadInt32(&count))
			runtime.Gosched()

			group.Done()
		}()
		fmt.Println(count)
		fmt.Println("goroutines: ", runtime.NumGoroutine())

	}

	group.Wait()
	fmt.Println(count)
	fmt.Println("goroutines: ", runtime.NumGoroutine())

}
