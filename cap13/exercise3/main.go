package main

import "fmt"

func main() {
	defer defered()
	fmt.Println("after")
}

func defered() {
	fmt.Println("the last one")
}
