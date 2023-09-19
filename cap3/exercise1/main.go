package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	a := fmt.Sprint(x, y, z)

	fmt.Printf("%v", a)

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}
