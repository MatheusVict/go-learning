package main

import "fmt"

func main() {

	sliceOfSlice := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println(sliceOfSlice)
	fmt.Println(sliceOfSlice[1])
	fmt.Println(sliceOfSlice[0][2])

	sliceSuper := [][][][]int{
		[][][]int{
			[][]int{
				[]int{1, 2, 3, 4},
			},
			[][]int{
				[]int{4, 5, 6, 6},
			},
		},

		[][][]int{
			[][]int{
				[]int{9, 8, 7},
			},
		},
	}

	fmt.Println(sliceSuper[0][1][0][1])

}
