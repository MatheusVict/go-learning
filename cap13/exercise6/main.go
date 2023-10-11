package main

import "fmt"

func main() {

	x := func() string {
		return "your mon"
	}

	fmt.Println(x())

}
