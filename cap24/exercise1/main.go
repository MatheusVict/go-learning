package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {

	p1 := person{Last: "fasd", First: "fjasdlçk", Sayings: []string{"çdklfajs", "fjdkla"}}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}
