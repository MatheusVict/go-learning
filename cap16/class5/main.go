package main

import (
	"fmt"
	"sort"
)

type car struct {
	name    string
	horses  int
	consumo int
}

type orderByHorse []car

type orderByConsumo []car

type orderByPostOwnerLucro []car

func (horse orderByHorse) Len() int {
	return len(horse)
}

func (horse orderByHorse) Less(i, j int) bool {
	return horse[i].horses < horse[j].horses
}

func (horse orderByHorse) Swap(i, j int) {
	horse[i], horse[j] = horse[j], horse[i]
}

func (horse orderByConsumo) Len() int {
	return len(horse)
}

func (horse orderByConsumo) Less(i, j int) bool {
	return horse[i].consumo > horse[j].consumo
}

func (horse orderByConsumo) Swap(i, j int) {
	horse[i], horse[j] = horse[j], horse[i]
}

func (horse orderByPostOwnerLucro) Len() int {
	return len(horse)
}

func (horse orderByPostOwnerLucro) Less(i, j int) bool {
	return horse[i].consumo > horse[j].consumo
}

func (horse orderByPostOwnerLucro) Swap(i, j int) {
	horse[i], horse[j] = horse[j], horse[i]
}

func main() {
	cars := []car{
		{name: "Omega", horses: 600, consumo: 1},
		{name: "Opala", horses: 700, consumo: 0},
		{name: "Astra", horses: 200, consumo: 9},
		{name: "Civic", horses: 90, consumo: 15},
	}

	fmt.Println(cars)
	// pass a personal sort
	sort.Sort(orderByHorse(cars))

	fmt.Println(cars)

	sort.Sort(orderByConsumo(cars))
	fmt.Println(cars)

	sort.Sort(orderByPostOwnerLucro(cars))
	fmt.Println(cars)
}
