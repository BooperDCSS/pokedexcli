package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// used for the map and mapb funtions

func (c *Client) ListLocationAreas(pageURL *string) (RespShallowLocAreas, error) {
	url := baseURL
	if pageURL != nil {
		url = *pageURL
	}

	var ListOfLocations RespShallowLocAreas

	if cachedData, exists := c.clientCache.Get(url); exists {
		if err := json.Unmarshal(cachedData, &ListOfLocations); err != nil {
			return RespShallowLocAreas{}, err
		}
		fmt.Println("***Data retrieved from the cache***")
		return ListOfLocations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocAreas{}, err
	}

	resp, err := c.httpClient.Do(req) // note that the httpClient is a ref to the Client's struct field
	if err != nil {
		return RespShallowLocAreas{}, err
	}

	if resp.StatusCode > 299 {
		return RespShallowLocAreas{}, fmt.Errorf(
			"Bad status code: %d - map process terminated",
			resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocAreas{}, err
	}

	if err := json.Unmarshal(data, &ListOfLocations); err != nil {
		return RespShallowLocAreas{}, err
	}

	c.clientCache.Add(url, data) // cache the data here using the url that obtained the data

	return ListOfLocations, nil
}
