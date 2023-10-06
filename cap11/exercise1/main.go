package main

import "fmt"

type person struct {
	name     string
	lastname string
	tastes   []string
}

func main() {

	x := person{name: "maraco", lastname: "batata", tastes: []string{"feijão", "arroz"}}
	y := person{name: "maraco", lastname: "batata", tastes: []string{"feijão", "arroz"}}

	for i, i2 := range x.tastes {
		fmt.Println(i, i2)
	}

	for i, i2 := range y.tastes {
		fmt.Println(i, i2)
	}

}
