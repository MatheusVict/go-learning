package main

import "fmt"

func main() {
	x, y := str()
	fmt.Println(fun(), x, y)
}

func fun() int {
	return 0
}

func str() (int, string) {
	return 10, "fsafa"
}
