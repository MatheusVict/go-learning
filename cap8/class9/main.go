package main

import "fmt"

func main() {
	friends := map[string]int{
		"me":  456131,
		"he":  43512153,
		"she": 615351,
	}

	fmt.Println(friends)
	fmt.Println(friends["me"])

	friends["gopher"] = 4444

	fmt.Println(friends)

	if perhaps, ok := friends["ghost"]; !ok {
		fmt.Println("it don't", ok)
	} else {
		fmt.Println(perhaps)
	}
	if perhaps, ok := friends["gopher"]; !ok {
		fmt.Println("it don't", ok)
	} else {
		fmt.Println(perhaps)
	}

}
