package main

import "fmt"

func main() {
	x := struct {
		name string
		age  int
	}{
		name: "batata",
		age:  50,
	}

	fmt.Println(x)
}
