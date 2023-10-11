package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}

func changeMe(p *person) {
	(*p).name = "hello"
	(*p).lastname = "it's me"
	(*p).age = 15

	fmt.Println((*p).name)

}

func main() {

	x := person{name: "fjklçads", lastname: "fdjaslç", age: 0}
	changeMe(&x)
	fmt.Println(x)

}
