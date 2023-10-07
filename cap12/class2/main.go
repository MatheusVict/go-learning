package main

import "fmt"

func main() {
	s := []int{
		10,
		20,
		3,
		04,
		56,
	}

	fmt.Println(sum(s...))
}

func sum(x ...int) int {
	value := 0
	for _, v := range x {
		value += v
	}
	return value
}
