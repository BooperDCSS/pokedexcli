package main

import (
	"time"

	"github.com/BooperDCSS/pokedexcli/internal/pokeapi"
	"github.com/BooperDCSS/pokedexcli/internal/pokecache"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second) // this client times out after 5 seconds

	// the uninitialized config fields for the next and prev location URLS is intentionally blank
	// Go sets uninitialized struct fields to their zero values for you
	// that's nil for pointers, which is what we want
	conf := &config{
		Commands:      getCommands(),
		pokeapiCache:  pokecache.NewCache(20 * time.Second), // only exists in the config; memory addressable!
		pokeapiClient: pokeClient,
	}
	replInput(conf)
}
