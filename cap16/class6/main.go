package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "hello"
	sb, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}

	wrongPassword := "dfjkçlas"

	fmt.Println(string(sb))
	erro := bcrypt.CompareHashAndPassword(sb, []byte(wrongPassword))
	fmt.Println(erro)
}

//func CompareHashAndPassword(hashedPassword, password []byte) error
