package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	converge := make(chan int)

	go send(even, odd)
	go receive(even, odd, converge)

	for v := range converge {
		fmt.Println("Received value: ", v)
	}

}

func send(even, odd chan int) {
	x := 100
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd, converge chan int) {
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for v := range even {
			converge <- v
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range odd {
			converge <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(converge)
}
