package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to Treasure Island")
	fmt.Println("Your mission is to find the treasure")
	time.Sleep(3 * time.Second)
	clear()
	crossroads()
	clear()
	lake()
	clear()
	house()
	clear()
	fmt.Println("You found the treasure!  You Win!")
}

func crossroads() {
	var choice string
	fmt.Println("You are at a cross road.  Where do you want to go? (Left or Right) ")
	fmt.Scanln(&choice)
	if choice == "" {
		fmt.Println("You failed to move in a valid direction.  You grow old, and die in this spot.")
		os.Exit(1)
	}
	choice = strings.ToLower(choice)
	if choice == "left" || choice == "l" {
		fmt.Println("You fell into a hole.")
		os.Exit(1)
	}
	if choice == "right" || choice == "r" {
		return
	}
	fmt.Println("You failed to follow instructions and are eaten by a Gru.  Game Over.")
	os.Exit(1)
}

func lake() {
	var choice string
	fmt.Println("You have come to a lake.  There is an island in the middle of the lake.  (Wait or Swim) ")
	fmt.Scanln(&choice)
	if choice == "" {
		fmt.Println("While you stand there doing nothing, a meteor falls from the sky obliterating you.  You Lose")
		os.Exit(1)
	}
	choice = strings.ToLower(choice)
	if choice == "swim" || choice == "s" {
		fmt.Println("You get attacked by an angry trout.  Game Over.")
		os.Exit(1)
	}
	if choice == "wait" || choice == "w" {
		return
	}
	fmt.Println("You failed to follow instructions and are eaten by a Gru.  Game Over.")
	os.Exit(1)
}

func house() {
	var choice string
	fmt.Println("You arrive at the island unharmed.  There is a house with three doors.  One Red, One Yellow, One Blue. (Which color do you choose) ")
	fmt.Scanln(&choice)
	if choice == "" {
		fmt.Println("While standing dumbly in front of the house a band of vikings cross the lake.  You die in a brutal and violent fashion.  Your last sight is the house being burned to the ground.")
		os.Exit(1)
	}
	choice = strings.ToLower(choice)
	if choice == "red" || choice == "r" {
		fmt.Println("You enter a room full of beasts.  Game Over.")
		os.Exit(1)
	}
	if choice == "blue" || choice == "b" {
		fmt.Println("You enter a room full of beasts.  Game Over.")
		os.Exit(1)
	}
	if choice == "yellow" || choice == "y" {
		return
	}
	fmt.Println("You failed to follow instructions and are eaten by a Gru.  Game Over.")
	os.Exit(1)
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
