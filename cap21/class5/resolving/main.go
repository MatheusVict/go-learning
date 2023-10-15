package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go sendNumbers(even, odd, quit)

	receive(even, odd, quit)
}

func receive(even chan int, odd chan int, quit chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even number: ", v)
		case v := <-odd:
			fmt.Println("ODD number: ", v)
		case v, ok := <-quit:
			if !ok {
				fmt.Println("It's bad", v)
			} else {
				fmt.Println("Turning off", v)

			}
			return
		}
	}
}

func sendNumbers(even chan int, odd chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
	quit <- true
}
