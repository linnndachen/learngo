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
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please give me 2 numbers.")
	}

	num1, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		fmt.Printf("%s is not a number", os.Args[1])
	}

	num2, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		fmt.Printf("%s is not a number", os.Args[2])
	}

	if num1 <= 0 || num2 <= 0 {
		fmt.Println("Please enter only positive number")
	}

	rand.Seed(time.Now().UnixNano())
	min := num1
	if num1 < num2 {
		min = num2
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(min) + 1

		if n == num1 || n == num2 {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")

}
