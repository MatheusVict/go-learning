package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("itsme.texte")
	if err != nil {
		fmt.Println("Err happened: ", err)
		log.Println("Err happened: ", err) // log: 2023/10/16 11:08:53 Err happened:  open itsme.texte: no such file or directory

	}
}
