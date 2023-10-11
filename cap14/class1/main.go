package main

import "fmt"

func main() {
	z := 10

	y := &z

	//0xc000094000
	fmt.Println(&z, y)
	// show what has in a address
	fmt.Println(*y)

	fmt.Println(*&y)
	// *int == pointer
	fmt.Printf("%T", y)
}
