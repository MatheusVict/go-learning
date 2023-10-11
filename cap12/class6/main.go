package main

import "fmt"

func main() {
	x := 10

	// Anonymity function
	func(x int) {
		fmt.Println(x * 546481)
	}(x)

}
