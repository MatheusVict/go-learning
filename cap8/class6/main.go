package main

import "fmt"

func main() {
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
	slice = append(slice, 11)
	slice = append(slice, 12)
	slice = append(slice, 13)
	slice = append(slice, 14)
	slice = append(slice, 15)

	fmt.Println(slice, len(slice), cap(slice))
}
