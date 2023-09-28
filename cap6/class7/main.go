package main

import "fmt"

var y interface{}

func main() {
	x := 10

	switch {
	case x < 5:
		fmt.Println("hello")
		break
	case x == 10:
		fmt.Println("yea")
	}

	switch true {
	case x < 5:
		fmt.Println("fsad")
		break
	case x == 10:
		fmt.Println("fdasasfafdas")
	}

	switch true {
	case x == 10:
		fmt.Println("fasdddddddddddddddddddddddddddd")
		fallthrough
	case x < 5:
		fmt.Println("afdsssssssssssssssssssssssssssssss")
	}

	switch true {
	case (x < 5), (x%2 == 0):
		fmt.Println("par")
	case x == 10, x > 9:
		fmt.Println("biggerr than")
	}

	switch x {
	case x + 5:
		fmt.Println("aaa")
	}

	y = true

	switch y.(type) {
	case int:
		println("int")
	case float64:
		println("float")
	case string:
		println("string")
	case bool:
		print("bool")
	}

	switch z := 0; {
	case z == 1:
		println("int")
	case z == 2:
		println("float")
	case z == 3:
		println("string")
	case z == 4:
		print("bool")
	}
}
