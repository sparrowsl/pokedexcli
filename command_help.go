package main

import "fmt"

func callbackHelp() {
	fmt.Println("Welcome to the Pokedex help menu")
	fmt.Println("Here are your available commands:")

	for _, command := range getCommands() {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Println()
}
