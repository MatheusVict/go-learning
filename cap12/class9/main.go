package main

import "fmt"

func main() {
	x := sumOnlyPar(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(x)
}

func sum(x ...int) int {
	n := 0
	for _, i := range x {
		n += i
	}
	return n
}

func sumOnlyPar(f func(x ...int) int, y ...int) int {
	var slice []int

	for _, i := range y {
		if i%2 == 0 {
			slice = append(slice, i)
		}
	}
	total := f(slice...)
	return total
}
