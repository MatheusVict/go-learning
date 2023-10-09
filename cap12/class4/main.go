package main

import "fmt"

type person struct {
	name string
	age  int
}

// quase uma funçao de extensão do kotlin
func (p person) hello() {
	fmt.Println(p.name, " say hello")
}

func main() {
	mauricio := person{name: "Mauricio", age: 20}
	mauricio.hello()
}
