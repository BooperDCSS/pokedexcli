package main

import (
	"fmt"

	"github.com/BooperDCSS/pokedexcli/internal/config"
)

func commandHelp(conf *config.Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, command := range conf.Commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}
