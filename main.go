package main

import (
	"github.com/BooperDCSS/pokedexcli/internal/config"
)

func main() {
	locationNext := "https://pokeapi.co/api/v2/location-area/"
	var locationPrevious *string

	var conf = &config.Config{
		Commands: getCommands(),
		Next:     &locationNext,
		Previous: locationPrevious,
	}
	replInput(conf)
}
