package main

import "fmt"

func main() {

	channel := make(chan int)

	// They need to run as concurrent
	go send(channel)
	receiver(channel)

}

/**
* <-chan to get
* chan-> to send
 */

func receiver(channel <-chan int) {
	fmt.Println("Received value:", <-channel)
}

func send(channel chan<- int) {
	channel <- 42
}
