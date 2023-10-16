package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("itsme.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	defer fmt.Println("Function ran")

	r := strings.NewReader("james bond")

	io.Copy(f, r)
}
