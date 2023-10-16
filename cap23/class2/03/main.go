package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	_, err := os.Open("no-file.txt")
	if err != nil {
		// log.Fatalln(err) //Fatalln is equivalent to Println() followed by a call to os.Exit(1).
		log.Panicln(err) // Panicln is equivalent to Println() followed by a call to panic()
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}
