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

		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
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
		"mapb": {
			name:        "mapb",
			description: "List all the previous poke maps",
			callback:    callbackMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)

	return words
}
