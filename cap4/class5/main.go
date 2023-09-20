package main

import "fmt"

func main() {
	s := "ascíí éáó, ¨s"
	sb := []byte(s)

	for _, v := range sb {

		fmt.Printf("%v - %T - %#U - %#X\n", v, v, v, v)
	}
	for _, v := range s {

		fmt.Printf("%v - %T - %#U - %#X\n", v, v, v, v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#X\n", s[i], s[i], s[i], s[i])
	}

}

/*
* %v %T
* Raw string literals
* Converter to slice of bytes: []byte(x)
* %#U %#X
 */
