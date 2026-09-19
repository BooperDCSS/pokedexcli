package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BooperDCSS/pokedexcli/internal/config"
)

func replInput(conf *config.Config) {
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

		commandRequest := cleanedInput[0]

		request, exists := conf.Commands[commandRequest]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := request.Callback(conf)
		if err != nil {
			fmt.Println(err)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error encountered: %v", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
