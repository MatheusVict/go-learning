package main

import "fmt"

func main() {
	x := each()

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

}

func each() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
