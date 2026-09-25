package main

import (
	"errors"
	"fmt"
)

func commandMapForward(conf *config, optional ...string) error {

	locationResp, err := conf.pokeapiClient.ListLocationAreas(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.currentLocations = locationResp
	conf.nextLocationsURL = locationResp.Next // this makes it so conf.next and loc.Next point to the same mem location
	conf.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapBack(conf *config, optional ...string) error {
	if conf.previousLocationsURL == nil {
		return errors.New("you're on the first page") // bubbles up through the replInput function
	}
	locationResp, err := conf.pokeapiClient.ListLocationAreas(conf.previousLocationsURL)
	if err != nil {
		return err
	}

	conf.currentLocations = locationResp
	conf.nextLocationsURL = locationResp.Next
	conf.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
