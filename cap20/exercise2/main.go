package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}

type humans interface {
	talk()
}

func (p *person) talk() {
	fmt.Println("Hello ", (*p).name)
}

func toSaySomething(human humans) {
	human.talk()
}

func main() {

	pers := person{name: "yeaeaah", lastname: "Idk", age: 19}
	toSaySomething(&pers)
	pers.talk()    // shortcut for (&pers).talk()
	(&pers).talk() // the correct one
}
