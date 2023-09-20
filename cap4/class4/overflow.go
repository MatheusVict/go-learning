package main

import "fmt"

func main() {
	var i uint16
	i = 65535
	fmt.Println(i)
	i++ // when you pass from limit it becomes 0
	println(i)
}
