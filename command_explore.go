package main

import (
	"fmt"
)


func commandExplore(conf *config, locationName ...string) error {
	// dig into the deeper data we're going to get from our request
	// find all the pokemon in that area
	// return their names
	// will need to use the ListPokemonInArea function...

	if len(locationName) == 0 {
		fmt.Println("'explore' requires you enter a location name.")
		return nil
	}

	if len(locationName) > 1 {
		fmt.Println("Please select only one location from the current list.")
		return nil
	}
	
	deepResp, err := conf.pokeapiClient.GetDeepLocationData(locationName[0])
	if err != nil {
		return err
	}

	fmt.Printf("Explore %s...\n", deepResp.Name)
	for _, entry := range deepResp.PokemonEncounters {
		fmt.Printf("- %s\n", entry.Pokemon.Name)
	} 
	
	return nil
}