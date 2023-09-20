package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "é"
	b := "e"
	c := "!@@#$%¨&*()_+"

	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a) // so i can see in byte form
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v\n", d, e, f)

	x := 10
	y := 10.0 // float is better when we use int64

	fmt.Printf("%v, %T", x, x)
	fmt.Printf("%v, %T\n", y, y)

	fmt.Println(runtime.GOOS)   // os
	fmt.Println(runtime.GOARCH) // architecture
}
