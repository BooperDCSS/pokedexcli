package main

import "fmt"


func commandExplore(conf *config, locationName ...string) error {
	// dig into the deeper data we're going to get from our request
	// find all the pokemon in that area
	// return their names
	// will need to use the ListPokemonInArea function...
	fmt.Println("Exploring...")
	for _, loc := range conf.currentLocations.Results {
		fmt.Printf("name: %s - URL: %s\n", loc.Name, loc.URL)
	}
	return nil
}