package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}

func (p person) getPeople() {
	fmt.Println("Hello ", p.name, " ", p.lastname, " are you ", p.age)
}

func main() {

	people := person{name: "Ahsoka", age: 35, lastname: "Tano"}

	people.getPeople()

}
