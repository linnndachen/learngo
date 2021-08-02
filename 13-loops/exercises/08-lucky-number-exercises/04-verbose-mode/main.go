// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?

(Provide -v flag to see the picked numbers.)
`
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var verbose bool

	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number")
		return
	}

	rand.Seed(time.Now().UnixNano())
	for turns := 1; turns <= maxTurns; turns++ {
		n := rand.Intn(guess) + 1

		if verbose {
			fmt.Printf("%d", n)
		}

		if n == guess {
			if turns == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU WON!")
			}
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")

}
