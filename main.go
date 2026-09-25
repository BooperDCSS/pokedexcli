package main

import (
	"time"

	"github.com/BooperDCSS/pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	// the uninitialized config fields for the next and prev location URLS is intentionally blank
	// Go sets uninitialized struct fields to their zero values for you
	// that's nil for pointers, which is what we want
	conf := &config{
		Commands:      getCommands(),
		pokeapiClient: pokeClient,
	}
	replInput(conf)
}
