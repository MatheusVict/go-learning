package main

import (
	"fmt"
	"math"
)

type square struct {
	side int
}

type circle struct {
	raio int
}

func (c circle) area() float64 {
	return math.Pi * float64(2) * float64(c.raio)
}

func (s square) area() float64 {
	return float64(s.side) * float64(s.side)
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {

	squa := square{side: 15}
	cir := circle{raio: 46}

	fmt.Println(info(squa))
	fmt.Println(info(cir))
}
