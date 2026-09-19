package main

import (
	"fmt"
	"os"

	"github.com/BooperDCSS/pokedexcli/internal/config"
)

func commandExit(conf *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
