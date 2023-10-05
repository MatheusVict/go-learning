package main

import "fmt"

func main() {
	firstSlice := []int{1, 2, 3, 4, 5}
	/*
		[1 2 3 4 5]
	*/
	fmt.Println(firstSlice)

	secondSlice := append(firstSlice[:2], firstSlice[4:]...)
	//[1 2 5]
	fmt.Println(secondSlice)
	//     !
	//[1 2 5 4 5]
	fmt.Println(firstSlice)
}
