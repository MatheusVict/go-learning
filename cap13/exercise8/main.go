package main

import "fmt"

func reco() func() int {
	return func() int {
		return 15
	}
}

func main() {

	x := reco()()
	fmt.Println(x)

}
