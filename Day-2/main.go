package main

import (
	"fmt"
)

func main() {
	var bill float64
	var people float64
	var tip float64

	fmt.Println("What is the total Bill? ")
	fmt.Scanln(&bill)

	fmt.Println("How many people are splitting it? ")
	fmt.Scanln(&people)

	fmt.Println("What Tip Percentage Are you leaving? ")
	fmt.Scanln(&tip)

	totalbill := ((bill * tip) / 100) + bill
	portion := totalbill / people

	fmt.Printf("Each person should pay $%.2f \n", portion)

}
