package main

import "fmt"

func main() {

	cs := make(chan int)
	go func() {
		cs <- 42
	}()

	fmt.Printf("cs\t%T\n", cs)
}
