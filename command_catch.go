package main

import (
	"fmt"
	"math/rand"
	"time"
)

// catch puts a pokemon in your conf.pokedex based on their Base Experience
// and a random dice d100 roll + 30... pretty sure some Pokemon can't be
// caught with this configuration

func commandCatch(conf *config, pokeName ...string) error {

	if len(pokeName) < 1 {
		fmt.Println("Use catch <pokemon_name> to try catching a Pokemon.")
		return nil
	}

	if len(pokeName) > 1 {
		fmt.Println("You can catch only one Pokemon at a time.")
		return nil
	}

	respPokemon, err := conf.pokeapiClient.GetPokeData(pokeName[0])
	if err != nil {
		return err
	}

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	catchPower := randSource.Intn(100) + 30

	fmt.Printf("Throwing a Pokeball at %s...\n", respPokemon.Name)

	if catchPower < respPokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", respPokemon.Name)
		return nil
	}

	conf.pokedex[respPokemon.Name] = respPokemon
	fmt.Printf("%s was caught!\n", respPokemon.Name)

	return nil
}
