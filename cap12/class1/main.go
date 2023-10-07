package main

import "fmt"

func main() {
	yea := thatsWhatSheSaid(5, 60)
	fmt.Println(yea)

	// multiple args
	fmt.Println(multiple("oi", 10, 10, 10))
}

// x int, y int -> x, y int
func thatsWhatSheSaid(x, y int) int {
	return y + x
}

func multiple(y string, x ...int) (int, int, string) {
	value := 0
	for i, i2 := range x {
		fmt.Println(i, i2)
		value += i + i2
	}

	return value, len(x), y
}
