package pokeapi

import (
	"net/http"
	"time"

	"github.com/frogonabike/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient  http.Client
	clientCache pokecache.Cache
}

func NewClient(timeout time.Duration, cacheTimeout time.Duration) Client {

	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		clientCache: pokecache.NewCache(cacheTimeout),
	}
}
