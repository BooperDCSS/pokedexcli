package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/BooperDCSS/pokedexcli/internal/pokeapi"
)

// TODO: this all needs to live with the httpClient, not in the config
// go create the cache in the client, then route everything via ListLocationAreas

func commandMapForward(conf *config) error {

	// cacheData, exists := conf.pokeapiCache.Get(*conf.nextLocationsURL)
	// if exists {
	// 	var cachedLocation pokeapi.RespShallowLocAreas

	// 	if err := json.Unmarshal(cacheData, &cachedLocation); err != nil {
	// 		return err
	// 	}

	// 	conf.nextLocationsURL = cachedLocation.Next
	// 	conf.previousLocationsURL = cachedLocation.Previous

	// 	fmt.Println("Returning data from cache")
	// 	for _, loc := range cachedLocation.Results {
	// 		fmt.Println(loc.Name)
	// 	}
	// 	return nil
	// }

	// if no cached data, retrieve it from the client
	locationResp, err := conf.pokeapiClient.ListLocationAreas(conf.nextLocationsURL)

	if err != nil {
		return err
	}

	storeData, err := json.Marshal(locationResp) // can't think of a better way to do this
	if err != nil {
		return err
	}


	conf.nextLocationsURL = locationResp.Next         // this makes it so conf.next and loc.Next point to the same mem location
	conf.previousLocationsURL = locationResp.Previous // safe because it doesn't mutate state, just reassigns value

	conf.pokeapiCache.Add(*conf.nextLocationsURL, storeData) // store the raw bytes

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapBack(conf *config) error {
	if conf.previousLocationsURL == nil {
		return errors.New("you're on the first page") // bubbles up through the replInput function
	}

	cacheData, exists := conf.pokeapiCache.Get(*conf.previousLocationsURL)
	if exists {
		var cachedLocation pokeapi.RespShallowLocAreas

		if err := json.Unmarshal(cacheData, &cachedLocation); err != nil {
			return err
		}

		conf.nextLocationsURL = cachedLocation.Next
		conf.previousLocationsURL = cachedLocation.Previous

		fmt.Println("Returning data from cache")
		for _, loc := range cachedLocation.Results {
			fmt.Println(loc.Name)
		}
		return nil
	}

	locationResp, err := conf.pokeapiClient.ListLocationAreas(conf.previousLocationsURL)
	if err != nil {
		return err
	}
	storeData, err := json.Marshal(locationResp)
	if err != nil {
		return err
	}
	conf.pokeapiCache.Add(*conf.previousLocationsURL, storeData)

	conf.nextLocationsURL = locationResp.Next
	conf.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
