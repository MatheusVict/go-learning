package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 6, 98, 7, 4, 61, 1, 52, 6, 8, 9, 5, 23, 65}
	xs := []string{"arroz", "feijão", "batata", "oque falta?"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
