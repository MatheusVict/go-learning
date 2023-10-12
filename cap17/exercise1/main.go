package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {

	u1 := User{First: "me", Age: 19}
	u2 := User{First: "he", Age: 29}
	u3 := User{First: "she", Age: 39}
	u4 := User{First: "they", Age: 49}

	users := []User{u1, u2, u3, u4}

	fmt.Println(users)

	sliceOfBytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	println(string(sliceOfBytes))

}
