package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Name  string `json:"Name"` // tags
	Order string `json:"Order"`
}

func main() {

	sb := []byte(`{"Name": "Platypus", "Order": "Monotremata"}`)

	var inf info
	err := json.Unmarshal(sb, &inf)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(inf)

	fmt.Println(inf.Order)
	fmt.Println(inf.Name)
}
