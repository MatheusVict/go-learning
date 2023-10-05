package main

import "fmt"

func main() {

	anything := map[int]string{
		0: "arrroz",
		1: "that's what she said",
		2: "yeaaah",
	}

	for key, value := range anything {
		fmt.Println(key, value)
	}
}
