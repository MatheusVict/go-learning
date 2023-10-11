package main

import "fmt"

func main() {
	x := 0
	receiveAValue(x)
	receiveAPointer(&x)

}

func receiveAValue(x int) {
	x++
	fmt.Println(x)

}

func receiveAPointer(x *int) {
	*x++
	fmt.Println(*x)
}
