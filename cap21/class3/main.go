package main

import (
	"fmt"
	"sync"
)

var group = sync.WaitGroup{}

func main() {
	channel := make(chan int)

	group.Add(2)

	go myLoop(10, channel)
	go prints(channel)

	for v := range channel {
		fmt.Println("Received from channel: ", v)
	}
	group.Wait()
}

func myLoop(total int, send chan<- int) {
	for i := 0; i < total; i++ {
		send <- i
	}
	// You must CLOSE
	group.Done()
	close(send)
}

func prints(receive <-chan int) {
	for v := range receive {
		fmt.Println("Received from: ", v)
	}
	group.Done()
}
