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

	bs, err := toJSON(p1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}

func toJSON(p1 interface{}) ([]byte, error) {
	bs, err := json.Marshal(p1)
	if err != nil {
		return bs, fmt.Errorf("error on %v, with type: %T Error: %v", bs, bs, err)
	}
	return bs, fmt.Errorf("error on %v, with type: %T Error: %v", string(bs), bs, err)
}
