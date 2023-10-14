package main

import (
	"fmt"
	"sync"
)

var group = sync.WaitGroup{}

func main() {
	goRoutines(2)

}

func goRoutines(countOfRoutines int) {
	group.Add(countOfRoutines)
	go f1()
	go f2()
	group.Wait()
}

func f1() {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	group.Done()
}
func f2() {
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	group.Done()
}
