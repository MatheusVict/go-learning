package main

import (
	"encoding/json"
	"os"
)

type infos struct {
	Name  string `json:"Name"` // tags
	Order string `json:"Order"`
}

func main() {

	hello := infos{Name: "açfjklsd", Order: "fçadjlsk"}

	enconder := json.NewEncoder(os.Stdout)
	enconder.Encode(hello)
}
