package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}

type driver struct {
	person
	bus    string
	salary float64
}

type developer struct {
	person
	languages string
	crazySize string
}

/*func (p person) hello() {
	fmt.Println("My name is ", p.name, p.lastname, " my age is ", p.age)
}*/

func (x driver) sayHello() {
	fmt.Println("Hello, I'm ", x.name, " I drive a ", x.bus, " and I receive ", x.salary)
}

func (x developer) sayHello() {
	fmt.Println("Hello dev ", x.name, x.lastname, " and I know this language ", x.languages, " and I'm ", x.crazySize)
}

type people interface {
	sayHello()
}

func human(p people) {
	p.sayHello()
	switch p.(type) {
	case driver:
		p.sayHello()
		fmt.Println(p.(driver).salary)

	case developer:
		fmt.Println(p.(developer).languages)
	}

}

func main() {
	/*drive:= driver{}
	drive.hello()*/

	drive := driver{
		salary: 12000,
		bus:    "mercedes",
		person: person{
			name:     "john",
			age:      15,
			lastname: "august",
		},
	}
	dev := developer{
		crazySize: "more than 8 thousand",
		languages: "Golang",
		person: person{
			name:     "bob",
			age:      54,
			lastname: "Rapaz",
		},
	}

	human(dev)
	fmt.Println(" ")
	human(drive)
}
