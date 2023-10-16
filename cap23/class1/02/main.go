package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string

	fmt.Println("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Favorite food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Favorite sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer3, answer2, answer1)

}
