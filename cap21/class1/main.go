package main

import "fmt"

func main() {
	/**
	* It buffers value
	* It has a rate limit     !
	 */
	channel := make(chan int, 1)

	channel <- 42

	fmt.Println(<-channel)

	routine()
}

func routine() {
	channel := make(chan int)
	go func() {
		channel <- 42
	}()

	fmt.Println(<-channel)
}
