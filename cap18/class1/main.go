package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)
	// add total of go routines
	fmt.Println(runtime.NumGoroutine())

	go func1()
	go func2()

	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())

	wg.Wait()
	// wait for them

}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("FUNC1: ", i)
		time.Sleep(20)
	}
	wg.Done()
	// It's Done

}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("FUNC2: ", i)
		time.Sleep(20)
	}
	wg.Done()
}
