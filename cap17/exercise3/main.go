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
	inf := infos{Name: "jafskldç", Order: "fjadklç"}

	json.NewEncoder(os.Stdout).Encode(inf)

}
