package main

import "fmt"

func main() {
	// execute as last
	defer fmt.Println("with")
	fmt.Println("without")

	fmt.Println()
}
