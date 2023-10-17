package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4)
	y := multiple(1, 3, 3)
	fmt.Println(x)
	fmt.Println(y)
}

func sum(x ...int) int {
	total := 0
	for _, i := range x {
		total += i
	}
	return total
}
func multiple(x ...int) int {
	total := 1
	for _, i := range x {
		total *= i
	}
	return total
}
