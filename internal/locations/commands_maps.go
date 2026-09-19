package locations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BooperDCSS/pokedexcli/internal/config"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap(conf *config.Config) error {

	if conf.Next == nil {
		fmt.Println("You are on the last page")
		return nil
	}
	var loc LocationArea // create a LocationArea for the API data

	res, err := http.Get(*conf.Next)
	if err != nil {
		return fmt.Errorf("Error with GET NextURL request: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&loc); err != nil {
		return fmt.Errorf("Error decoding JSON for NextURL: %w", err)
	}

	for i := 0; i < len(loc.Results); i++ {
		fmt.Printf("%s\n", loc.Results[i].Name)
	}

	conf.Next = loc.Next
	conf.Previous = loc.Previous
	return nil
}

func CommandMapBack(conf *config.Config) error {

	if conf.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	var loc LocationArea

	res, err := http.Get(*conf.Previous)
	if err != nil {
		return fmt.Errorf("Error with GET PreviousURL request: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&loc); err != nil {
		return fmt.Errorf("Error decoding JSON for PreviousURL: %w", err)
	}

	for i := 0; i < len(loc.Results); i++ {
		fmt.Printf("%s\n", loc.Results[i].Name)
	}

	conf.Previous = loc.Previous
	conf.Next = loc.Next
	return nil

}
