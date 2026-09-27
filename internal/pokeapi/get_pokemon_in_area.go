package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// used with the explore function

func (c *Client) GetDeepLocationData(locationInput string) (RespDeepLocationData, error) {
	url := baseURL + locationInput

	var deepInfo RespDeepLocationData

	if cachedData, exists := c.clientCache.Get(url); exists {
		if err := json.Unmarshal(cachedData, &deepInfo); err != nil {
			return RespDeepLocationData{}, err
		}
		fmt.Println("***Data retrieved from the cache***")
		return deepInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocationData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocationData{}, err
	}

	if resp.StatusCode > 299 {
		return RespDeepLocationData{}, fmt.Errorf(
			"Bad Status from Client: %d. That location does not exist.",
			resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocationData{}, err
	}

	if err := json.Unmarshal(data, &deepInfo); err != nil {
		return RespDeepLocationData{}, err
	}

	c.clientCache.Add(url, data)

	return deepInfo, nil
}
