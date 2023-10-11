package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(allArgs(1, 2, 3, 4, 5, 6))
	fmt.Println(hello(slice))
}

func allArgs(x ...int) int {
	n := 0
	for _, value := range x {
		n += value
	}
	return n
}

func hello(x []int) int {
	n := 0
	for _, value := range x {
		n += value
	}
	return n
}
