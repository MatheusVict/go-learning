package main

import "fmt"

func main() {

	slice := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for key, value := range slice {
		fmt.Println(key, value)
	}

	fmt.Printf("%T", slice)

}
