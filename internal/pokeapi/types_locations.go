package pokeapi

// response with "shallow" information about the location areas
// shallow because the Results struct only shows the name and URL, where we find more info

type RespShallowLocAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
