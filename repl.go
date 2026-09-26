package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BooperDCSS/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(conf *config, optional ...string) error
}

// we pass the API client to the config because we want to access interface data provided by the Poke site
// via the config itself; since we use the client to GET our location area values, and we get URLs
// via the pokeapi package ListLocations function, this config makes the inner workings of the
// client available in the main package

type config struct {
	Commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	currentLocations pokeapi.RespShallowLocAreas
	locationURLDict  map[string]string
}

func replInput(conf *config) {
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
		optionalParam := []string{}

		if len(cleanedInput) > 1 {
			optionalParam = cleanedInput[1:] // used to explore areas; passes a []string via the request
		}

		request, exists := conf.Commands[commandRequest]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := request.Callback(conf, optionalParam...)
		if err != nil {
			fmt.Println(err)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error encountered: %v", err)
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Print the next 20 Pokedex location-areas",
			Callback:    commandMapForward,
		},
		"mapb": {
			Name:        "mapb",
			Description: "mapb(ack) - Print the previous 20 Pokedex location-areas",
			Callback:    commandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore <area name> - prints all of the Pokemon found in that area",
			Callback:    commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
