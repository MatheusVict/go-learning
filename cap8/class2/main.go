package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(append(slice, 6))
	slice[3] = 45331

	fmt.Println(slice[3])

}
