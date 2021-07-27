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
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q\n"
	errPwd   = "Invalid password for %q\n"
	accessOK = "Access granted to %q\n"

	user = "jack"
	pin  = "1888"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	username, password := args[1], args[2]
	if username != user {
		fmt.Printf(errUser, username)
	} else if password != pin {
		fmt.Printf(errPwd, password)
	} else {
		fmt.Printf(accessOK, username)
	}
}
