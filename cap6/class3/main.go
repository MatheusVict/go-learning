package main

import "fmt"

func main() {

	for minute := 0; minute <= 12; minute++ {
		fmt.Println("MINUTE: ", minute)
		for x := 0; x < 60; x++ {
			fmt.Println("SECONDS: ", x)
		}
	}

}
