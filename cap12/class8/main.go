package main

import "fmt"

func main() {
	x := returnFunc()
	y := x(3)
	fmt.Println(y)
	fmt.Println(returnFunc()(3))
}

func returnFunc() func(i int) int {
	return func(i int) int {
		return i * 10

	}
}
