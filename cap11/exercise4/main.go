package main

import "fmt"

func main() {

	x := struct {
		name   string
		taste  string
		map1   map[string]string
		slice1 []string
	}{
		name:  "fasçdjlafsd",
		taste: "façsjlfa",
		//			[key]value
		map1:   map[string]string{"aflçsdkjfa": "fadskljçfdafadjskçld"},
		slice1: []string{"fajlçds", "fdçajlska"},
	}
	fmt.Println(x)

}
