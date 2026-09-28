package main

import "fmt"

func commandInspect(conf *config, pokemonName ...string) error {
	
	if len(pokemonName) < 1 {
		fmt.Println("Enter the name of at least one Pokemon you have captured.")
		return nil
	}

	if len(pokemonName) > 1 {
		fmt.Println("Enter the name of one captured Pokemon at a time.")
		return nil
	}

	thisPokemon, ok := conf.pokedex[pokemonName[0]]
	
	if !ok {
		fmt.Printf("You have not caught %s\n", pokemonName[0])
		return nil
	}

	fmt.Printf("Name: %s\n", thisPokemon.Name)
	fmt.Printf("Height: %d\n", thisPokemon.Height)
	fmt.Printf("Weight: %d\n", thisPokemon.Weight)
	fmt.Println("Stats:")

	for _, statEntry := range thisPokemon.Stats {
		fmt.Printf("    - %s: %d\n", statEntry.Stat.Name, statEntry.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeEntry := range thisPokemon.Types {
		fmt.Printf("    - %s\n", typeEntry.Type.Name)
	}
	return nil
}