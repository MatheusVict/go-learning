package main

import (
	"fmt"
	"runtime"
)

//GOOS=windows GOARCH=amd64 go build test.go

func main() {
	fmt.Println("Cruzada compilation, that's was linux/amd64, and now it's: ", runtime.GOARCH, runtime.GOOS)
}
