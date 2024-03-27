package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		// print the > if no text is entered
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		// switch on commands and perform functions based on the name
		switch command {
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)

	return words
}
