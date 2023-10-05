package main

import "fmt"

func main() {
	slice := []string{"oi", "io", "you"}

	for i, v := range slice {
		fmt.Println("in", i, "the value", v)
	}

	slice[1] = "hello"

	// range faz um range pelo slice
	for i, v := range slice {
		fmt.Println("in", i, "the value", v)
	}
}
