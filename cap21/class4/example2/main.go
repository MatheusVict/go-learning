package main

import "fmt"

func main() {

	channel := make(chan int)
	quit := make(chan int)

	go receiveQuit(channel, quit)
	sendToChannel(channel, quit)

}

func receiveQuit(channel chan int, quit chan int) {
	for i := 0; i < 500; i++ {
		fmt.Println("Received: ", <-channel)
	}
	quit <- 0
}
func sendToChannel(channel chan int, quit chan int) {
	anyThing := 10
	for {
		select {
		// send things to channel
		case channel <- anyThing:
			anyThing++
		// receive things from channel
		case <-quit:
			return
		}
	}
}
