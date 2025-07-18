package main

import (
	"fmt"
	"strings"
)

const (
	validAttemps = 2
	Username     = "wonder"
	Password     = "woman"
)

func main() {
	fmt.Println("WELCOME")
	//var declaration
	inputUser, inputPassword := "", ""
	//attemps := 0

	fmt.Println(inputUser, inputPassword)
	for attemps := 0; attemps <= validAttemps; attemps++ {
		fmt.Print("Insert user : ")
		fmt.Scanln(&inputUser)

		fmt.Print("Insert password: ")
		fmt.Scanln(&inputPassword)

		lowerUser := strings.ToLower(inputUser)
		lowerPassword := strings.ToLower(inputPassword)
		if lowerUser == Username && lowerPassword == Password {
			fmt.Println("Welcome to the Manchester United Bank")
		}
	}
	fmt.Println("Good look for the next one")
}
