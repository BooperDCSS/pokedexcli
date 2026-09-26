package main

import (
	"errors"
	"fmt"
)

func commandMapForward(conf *config, optional ...string) error {

	locationResp, err := conf.pokeapiClient.ListLocationAreas(conf.currentLocations.Next)
	if err != nil {
		return err
	}

	localDict := make(map[string]string)
	conf.currentLocations = locationResp

	// why save individual locations like this when we need access to the entire response for the Explore functionality?
	// conf.nextLocationsURL = locationResp.Next // this makes it so conf.next and loc.Next point to the same mem location
	// conf.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
		localDict[loc.Name] = loc.URL
	}

	conf.locationURLDict = localDict // used to make the explore functionality easy; maps names to URLs
	return nil
}

func commandMapBack(conf *config, optional ...string) error {
	if conf.currentLocations.Previous == nil {
		return errors.New("you're on the first page") // bubbles up through the replInput function
	}
	locationResp, err := conf.pokeapiClient.ListLocationAreas(conf.currentLocations.Previous)
	if err != nil {
		return err
	}

	localDict := make(map[string]string)
	conf.currentLocations = locationResp

	// we want the entire response in the config for Explore functionality anyway!
	// conf.nextLocationsURL = locationResp.Next
	// conf.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
		localDict[loc.Name] = loc.URL
	}

	conf.locationURLDict = localDict

	return nil
}
