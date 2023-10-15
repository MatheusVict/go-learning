package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 42
		close(channel)
	}()

	v, ok := <-channel

	// 42 false = is it a real channel value
	fmt.Println(v, ok)

	v, ok = <-channel

	// 0 true = is it a real channel value
	fmt.Println(v, ok)
}
