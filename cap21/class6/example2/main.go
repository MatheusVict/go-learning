package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := converge(work("First"), work("Second"))

	for i := 0; i < 15; i++ {
		fmt.Println(<-channel)
	}
}

func work(message string) chan string {
	channel := make(chan string)

	go func(m string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("FUNC %v say %v", m, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(message, channel)

	return channel
}

func converge(channelX, channelY chan string) chan string {
	newChannel := make(chan string)

	go func() {
		for {
			// new channel<- sets the return of <-channelX
			newChannel <- <-channelX
		}
	}()
	go func() {
		for {
			newChannel <- <-channelY
		}
	}()

	return newChannel
}
