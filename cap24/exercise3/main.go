package main

import (
	"errors"
	"fmt"
	"log"
)

type specialError struct {
	message string
	status  int
	err     error
}

func (s specialError) Error() string {
	return fmt.Sprintf("A error: %v ocurred on %v with status %v", s.message, s.err, s.status)
}

func main() {
	s := specialError{status: 500, message: "Internal error server", err: errors.New("error dad")}
	errorSla(s)

}

func errorSla(e error) {
	v := e.(specialError).message
	log.Println(v, e)
}
