package main

import "fmt"

func main() {
	x := [][]string{
		[]string{
			"Tú", "fofinha", "sleep",
		},
		[]string{
			"fdas", "gasdgas", "fasf Hobby",
		},
		[]string{
			"fads", "gdasga", "fadsf Hobby",
		},
	}

	for _, strings := range x {
		for _, v := range strings {
			fmt.Println(v, strings)
		}
	}

}
