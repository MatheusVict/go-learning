package main

import "fmt"

type client struct {
	name     string
	lastname string
	smoke    bool
}

func main() {
	client1 := client{name: "pepino", lastname: "arroz", smoke: false}
	client2 := client{name: "macarrão", lastname: "farofa", smoke: false}
	client3 := client{name: "jibalaia", lastname: "çasjlfkdk", smoke: false}

	fmt.Println(client2, "\n", client3, "\n", client1)
}
