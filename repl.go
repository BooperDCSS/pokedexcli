package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() { // blocks, waiting for input on the command
			break
		}
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue // prevents a panic if the user presses enter without typing anything
		}
		fmt.Printf("Your command was: %v\n", cleanedInput[0])
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error encountered: %v", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))

}
