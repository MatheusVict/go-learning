package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

// This method make this struct a valid error
func (n norgateMathError) Error() string {
	return fmt.Sprintf("A norgate math error occured: %v %v %v: ", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrrNorgateMath := fmt.Errorf("nor-gate math: square root of negative number: %v", f)

		return 0, norgateMathError{long: "50.1231 N", lat: "99.5562 W", err: ErrrNorgateMath}
	}
	return 42, nil
}
