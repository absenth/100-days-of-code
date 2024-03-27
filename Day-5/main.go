package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// this should take in a number of arugments
	// number of letters (upper and lower case)
	// number of digits
	// number of symbols
	var letter_choices string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTVUVWXYZ"
	var number_choices string = "1234567890"
	var symbol_choices string = "!@#$%^&*"
	var letter_pick int
	var number_pick int
	var symbol_pick int

	fmt.Println("Welcome to the passwordgen")
	fmt.Println("How many letters in your password? ")
	fmt.Scanln(&letter_pick)
	fmt.Println("How many numbers in your password? ")
	fmt.Scanln(&number_pick)
	fmt.Println("How many symbols in your password? ")
	fmt.Scanln(&symbol_pick)

	//genpass(letter_choices, number_choices, symbol_choices, letter_pick, number_pick, symbol_pick)

}

func genpass(letter_choices string, number_choices string, symbol_choices string, letter_pick int, number_pick int, symbol_pick int) {
	total_length := letter_pick + number_pick + symbol_pick

	for total_length > 0 {
		// do stuff...
		// genpass

	}
}
