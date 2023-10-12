package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"it's", "truth", "pretty", "lady"}
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)

	si := []int{5, 6, 0, 4, 31, 6154, 4, 0, 3, 24, 6, 4, 8, 7, 9}
	fmt.Println(si)
	sort.Ints(si)
	fmt.Println(si)
}
