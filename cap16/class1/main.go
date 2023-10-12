package main

import (
	"encoding/json"
	"fmt"
)

// To convert to json you need put everything in maiusculo
type Person struct {
	Name      string
	Lastname  string
	Age       int
	Profissao string
	BankCount float32
}

func main() {
	ahsoka := Person{Name: "Ahsoka", Lastname: "Tano", Profissao: "Not Jedi", Age: 34, BankCount: 15}
	ani := Person{Name: "Anakin", Lastname: "vader", Profissao: "Not Jedi", Age: 24, BankCount: 15}

	a, err := json.Marshal(ahsoka)
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(ani)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(a))
	fmt.Println(string(b))
}
