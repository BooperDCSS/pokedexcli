package pokeapi

import (
	"net/http"
	"time"

	"github.com/BooperDCSS/pokedexcli/internal/pokecache"
)

// creating a client because it offers more flexibility
// likely we will do more than just use the GET request
// we are wrapping an http.Client in a Client struct to further configure it

type Client struct {
	httpClient          http.Client
	clientCache         *pokecache.Cache
}

// a little strange, but we return a Client struct that contains an http.Client{}
// that http.Client{} can have its Timeout element configured via the NewClient(timeout) call

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		clientCache: pokecache.NewCache(cacheInterval),
	}
}
