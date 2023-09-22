package main

const (
	y = iota
	x = iota
	z = iota // or you can let it without iota like:
)

const (
	a = iota + 1000
	_
	c
)

func main() {

}
