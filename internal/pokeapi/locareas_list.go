package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// this looks like a lot, but it's just a GET and unmarshal function that returns the JSON data
// as a struct; it's used in the map functions so that the process of printing the maps
// and fetching the data is roughly separated, though the map functions will use the
// JSON struct to manipulate elements of the config struct, thereby changing the outcome of
// this function

func (c *Client) ListLocationAreas(pageURL *string) (RespShallowLocAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocAreas{}, err
	}

	resp, err := c.httpClient.Do(req) // note that the httpClient is a ref to the Client's struct field
	if err != nil {
		return RespShallowLocAreas{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocAreas{}, err
	}

	var loc RespShallowLocAreas
	if err := json.Unmarshal(data, &loc); err != nil {
		return RespShallowLocAreas{}, err
	}

	return loc, nil

}
