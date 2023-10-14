package main

import (
	"fmt"
	"runtime"
	"sync"
)

var group = sync.WaitGroup{}

func main() {

	count := 0

	numberOfGoroutines := 10

	group.Add(numberOfGoroutines)

	fmt.Println("goroutines: ", runtime.NumGoroutine())

	mutex := sync.Mutex{}

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			mutex.Lock()
			value := count

			runtime.Gosched()

			value++

			count = value

			mutex.Unlock()
			group.Done()
		}()
		fmt.Println(count)
		fmt.Println("goroutines: ", runtime.NumGoroutine())

	}

	group.Wait()
	fmt.Println(count)
	fmt.Println("goroutines: ", runtime.NumGoroutine())

}
