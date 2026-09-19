package main

import (
	"github.com/BooperDCSS/pokedexcli/internal/config"
	"github.com/BooperDCSS/pokedexcli/internal/locations"
)

func getCommands() map[string]config.CliCommand {
	return map[string]config.CliCommand{
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
			Callback:    locations.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "mapb(ack) - Print the previous 20 Pokedex location-areas",
			Callback:    locations.CommandMapBack,
		},
	}
}
