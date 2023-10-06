package main

import "fmt"

type person struct {
	name     string
	lastname string
	tastes   []string
}

func main() {

	x := person{name: "maraco", lastname: "batata", tastes: []string{"feijão", "arroz"}}
	y := person{name: "maraco", lastname: "oi", tastes: []string{"feijão", "arroz"}}

	mapdaddy := map[string]person{
		x.lastname: x,
		y.lastname: y,
	}

	for key, value := range mapdaddy {
		for _, taste := range value.tastes {
			fmt.Println(key, taste)
		}
	}

}
