package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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

		commandName := cleaned[0]
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.callback(cfg)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu with available commands",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the application",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "List all the poke maps",
			callback:    callbackMap,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)

	return words
}
