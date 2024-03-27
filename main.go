package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
