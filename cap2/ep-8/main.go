package main

import "fmt"

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%v\n", b)
	fmt.Printf("%v", x)
	b = hotdog(x)
}
