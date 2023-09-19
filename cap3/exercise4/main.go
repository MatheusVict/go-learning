package main

import "fmt"

type tu int

var x tu

func main() {
	println(x)
	fmt.Printf("%T\n", x)
	x = 42
	println(x)

}
