package main

import "fmt"

func main() {

	teastes := []string{"arroz", "feijão", "batata", "macarrão", "isso"}

	// start until
	fatia := teastes[0:2]
	fatia1 := teastes[0:2]
	fatia2 := teastes[2:len(teastes)]

	fmt.Println(fatia)
	fmt.Println(fatia1)
	fmt.Println(fatia2)

	for i := 0; i < len(teastes); i++ {
		fmt.Println(teastes[i])
	}

}
