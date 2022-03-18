package main

import "fmt"

func main() {
	for {
		var input string

		fmt.Print("Please input a command:")
		fmt.Scan(&input)

		switch input {
		case "help":
			fmt.Println("This is help command!")
		case "quit":
			return
		default:
			fmt.Println("This is other command!")
		}
	}
}