package main

import "fmt"

func main() {
	x := 2023
	for true {
		x--
		if x < 2004 {
			break
		} else {
			fmt.Println("I lived ", x)
		}
	}

}
