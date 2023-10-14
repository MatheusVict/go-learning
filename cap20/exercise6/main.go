package main

import (
	"fmt"
	"runtime"
)

func main() {

	//GOOS=windows GOARCH=amd64 go build main.go
	// go build *.go
	// ./FileName
	fmt.Println("My Os: ", runtime.GOOS, "My Arch: ", runtime.GOARCH)
}
