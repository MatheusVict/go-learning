package main

import "fmt"

type person struct {
	name string
	age  int
}

type proficional struct {
	person
	title   string
	salario int
}

func main() {
	person1 := person{name: "joão", age: 88}
	person2 := person{name: "felip", age: 12}
	person3 := person{name: "mané", age: 95}

	prossional1 := proficional{title: "padeiro", salario: 61532, person: person1}
	prossional2 := proficional{title: "çflasd", salario: 61532, person: person2}
	prossional3 := proficional{title: "çfjalds", salario: 61532, person: person3}

	fmt.Println(prossional1.person.name)
	fmt.Println(prossional2.name)
	fmt.Println(prossional3.age)
}
