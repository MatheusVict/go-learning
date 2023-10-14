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

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			value := count

			runtime.Gosched()

			value++

			count = value

			group.Done()
		}()
		fmt.Println(count)
		fmt.Println("goroutines: ", runtime.NumGoroutine())

	}

	group.Wait()
	fmt.Println(count)
	fmt.Println("goroutines: ", runtime.NumGoroutine())

}
