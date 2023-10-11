package main

import "fmt"

func main() {

	x := 10

	y := func(x int) int {
		return x * 65
	}

	fmt.Println(y(x))
}
