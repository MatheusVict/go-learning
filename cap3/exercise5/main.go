package main

import "fmt"

type tu int

var x tu

var y int

func main() {
	println(x)
	fmt.Printf("%T\n", x)
	x = 42
	println(x)

	y = int(x)

	fmt.Printf("%T", y)
}
