package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6, 7, 9}

	slice1 = append(slice1, 5, 6, 4, 7, 9)

	fmt.Println(slice1)

	// get only the elements inside
	slice1 = append(slice1, slice2...)

	fmt.Println(slice1)
}
