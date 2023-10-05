package main

import "fmt"

func main() {
	slice := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	firstToThird := slice[0:3]
	firthToLast := slice[4:len(slice)]
	secondeToSeven := slice[1:7]
	thirdToLast := slice[2 : len(slice)-1]

	fmt.Println(firstToThird, "\n", firthToLast, "\n", secondeToSeven, "\n", thirdToLast, "\n")
}
