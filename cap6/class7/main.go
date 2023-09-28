package main

import "fmt"

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
}
