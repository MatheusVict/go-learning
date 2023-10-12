package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	var us []user
	js := []byte(`[{"First":"me","Age":19},{"First":"he","Age":29},{"First":"she","Age":39},{"First":"they","Age":49}]`)

	fmt.Println(string(js))

	err := json.Unmarshal(js, &us)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(us)
}
