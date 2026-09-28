package main

import (
	"errors"
	"fmt"
)

func commandMapForward(conf *config, optional ...string) error {

	if len(optional) > 0 {
		fmt.Println("map doesn't take additional arguments")
	}

	locationResp, err := conf.pokeapiClient.ListLocationAreas(conf.currentLocations.Next)
	if err != nil {
		return err
	}

	conf.currentLocations = locationResp

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapBack(conf *config, optional ...string) error {

	if len(optional) > 0 {
		fmt.Println("mapb doesn't take additional arguments")
	}

	if conf.currentLocations.Previous == nil {
		return errors.New("you're on the first page") // bubbles up through the replInput function
	}
	locationResp, err := conf.pokeapiClient.ListLocationAreas(conf.currentLocations.Previous)
	if err != nil {
		return err
	}

	conf.currentLocations = locationResp

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
