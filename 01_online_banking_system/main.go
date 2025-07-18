package main

import (
	"fmt"
	"strings"
)

const (
	Attemps  = 2
	Username = "wonder"
	Password = "woman"
)

func main() {
	fmt.Println("WELCOME")

	inputUser, inputPassword := "", ""
	fmt.Print("Insert user : ")
	fmt.Scanln(&inputUser)

	fmt.Print("Insert password: ")
	fmt.Scanln(&inputPassword)

	lowerUser := strings.ToLower(inputUser)
	lowerPassword := strings.ToLower(inputPassword)

	fmt.Println(inputUser, inputPassword)
	if lowerUser == Username && lowerPassword == Password {
		fmt.Println("Welcome to the Manchester United Bank")
	} else {
		fmt.Println("Good look for the next time")
	}
}
