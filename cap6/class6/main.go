package main

import "fmt"

func main() {
	if x := 10; !(x > 100) {
		fmt.Println(x)
	}
}
