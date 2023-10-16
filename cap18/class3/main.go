package main

import (
	"fmt"
	"runtime"
	"sync"
)

var group = sync.WaitGroup{}

// go run -race test.go
func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())

	count := 0
	totalGoRoutines := 10

	group.Add(totalGoRoutines)

	var mu sync.Mutex

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			// lock
			mu.Lock()
			value := count

			//yield
			runtime.Gosched()

			value++
			count = value
			// unlock
			mu.Unlock()
			group.Done()
		}()
		fmt.Println("Goroutine: ", runtime.NumGoroutine())
	}

	group.Wait()
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("count: ", count)

}
