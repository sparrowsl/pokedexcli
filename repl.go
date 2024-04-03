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

		if err := command.callback(cfg, cleaned[1:]...); err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore [location_area]",
			description: "List the pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch [pokemon_name]",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)

	return words
}
