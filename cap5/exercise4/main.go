package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%b\n%d\n%#X\n", x, x, x)
	print("----------------\n")

	y := x << 1

	fmt.Printf("%b\n%d\n%#X", y, y, y)
}
