package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	funcs := 5

	go send(10, channel1)
	go another(funcs, channel1, channel2)

	for v := range channel2 {
		fmt.Println(v)
	}
}

func send(numberOfChannels int, channel chan int) {
	for i := 0; i < numberOfChannels; i++ {
		channel <- i
	}
	close(channel)
}

func another(funcs int, channel1 chan int, channel2 chan int) {
	var wg = sync.WaitGroup{}

	for i := 0; i < funcs; i++ {
		wg.Add(1)
		go func() {
			for v := range channel1 {
				channel2 <- work(v)
			}
			wg.Done()

		}()
	}
	wg.Wait()
	close(channel2)
	for v := range channel1 {
		go func(x int) {
			channel2 <- work(x)
		}(v)
	}

}

func work(number int) int {
	time.Sleep(time.Millisecond * 1000 /*time.Duration(rand.Intn(1e3))*/)
	return number
}
