package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {

	u1 := user{First: "me", Age: 19}
	u2 := user{First: "he", Age: 29}
	u3 := user{First: "she", Age: 39}
	u4 := user{First: "they", Age: 49}

	users := []user{u1, u2, u3, u4}

	fmt.Println(users)

	sliceOfBytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	println(string(sliceOfBytes))

}
