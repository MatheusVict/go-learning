package main

import "fmt"

func main() {

	channel1 := make(chan int)
	channel2 := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			channel1 <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			channel2 <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		// could receive for many channels
		select {
		case v := <-channel1:
			fmt.Println("Channel 1: ", v)
		case v := <-channel2:
			fmt.Println("Channel 2: ", v)
		}
	}

}
