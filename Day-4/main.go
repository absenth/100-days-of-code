package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var choices = [3]string{"rock", "paper", "scissors"}
	var player_pick string

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(choices))
	computer_pick := choices[randomIndex]

	clear()
	fmt.Println("Pick Rock, Paper, or Scissors")
	fmt.Scanln(&player_pick)
	player_pick = strings.ToLower(player_pick)

	fmt.Printf("Player Picked %s and Computer picked %s \n", player_pick, computer_pick)

	if computer_pick == "rock" && player_pick == "rock" {
		fmt.Println("You both picked Rock.  You tie!")
	}

	if computer_pick == "rock" && player_pick == "paper" {
		fmt.Println("Your paper covers the computer's Rock.  You Win!")
	}

	if computer_pick == "rock" && player_pick == "scissors" {
		fmt.Println("Computer's rock smashes your scissors.  You Lose!")
	}

	if computer_pick == "paper" && player_pick == "rock" {
		fmt.Println("Computer's paper covers your rock.  You lose!")
	}

	if computer_pick == "paper" && player_pick == "paper" {
		fmt.Println("You both picked paper.  You tie!")
	}

	if computer_pick == "paper" && player_pick == "scissors" {
		fmt.Println("Computer's paper is cut to ribbons by your scissors.  You win!")
	}

	if computer_pick == "scissors" && player_pick == "rock" {
		fmt.Println("Computer's scissors are smashed to bits by your rock.  You win!")
	}

	if computer_pick == "scissors" && player_pick == "paper" {
		fmt.Println("Computer's scissors cut your paper to ribbons.  You lose!")
	}

	if computer_pick == "scissors" && player_pick == "scissors" {
		fmt.Println("You both picked scissors.  You tie!")
	}

}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
