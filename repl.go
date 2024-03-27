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

		if scanner.Text() == "exit" {
			break
		}
		fmt.Println(scanner.Text())
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	words := strings.Fields(lowered)

	return words
}
