package main

import (
	"fmt"
)

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, command := range conf.Commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}
