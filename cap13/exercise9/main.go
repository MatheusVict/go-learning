package main

import "fmt"

func less(x ...int) int {
	hello := 0
	for _, v := range x {
		hello += v
	}
	return hello
}

func callEachOther(function func(x ...int) int) int {
	x := []int{1, 2, 3}
	return function(x...)

}

func main() {

	fmt.Println(callEachOther(less))

}
