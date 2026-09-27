package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// for the catch command; stuck with Unmarshal so I can use data in cache

func (c *Client) GetPokeData(pokemonName string) (RespPokeStats, error) {
	url := statsURL + pokemonName

	var PokemonStats RespPokeStats

	if cachedData, exists := c.clientCache.Get(url); exists {
		if err := json.Unmarshal(cachedData, &PokemonStats); err != nil {
			return RespPokeStats{}, err
		}
		fmt.Println("***Data retrieved from cache***")
		return PokemonStats, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokeStats{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokeStats{}, err
	}

	if resp.StatusCode > 299 {
		return RespPokeStats{}, fmt.Errorf(
			"Bad Status from Client: %d - no Pokemon with that name found",
			resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokeStats{}, err
	}
	
	if err := json.Unmarshal(data, &PokemonStats); err != nil {
		return RespPokeStats{}, err
	}

	c.clientCache.Add(url, data)

	return PokemonStats, nil
}
