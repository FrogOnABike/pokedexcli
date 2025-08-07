package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (respPokeonInfo, error) {
	pageURL := baseURL + "/pokemon/" + pokemonName

	cacheData, exists := c.clientCache.Get(pageURL)
	if exists {
		fmt.Println("***Retrieving from cache***")
		pokemonDetail := respPokeonInfo{}

		if err := json.Unmarshal(cacheData, &pokemonDetail); err != nil {
			return respPokeonInfo{}, err
		}
		return pokemonDetail, nil
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return respPokeonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return respPokeonInfo{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return respPokeonInfo{}, err
	}

	pokemonDetail := respPokeonInfo{}
	c.clientCache.Add(pageURL, body)

	if err := json.Unmarshal(body, &pokemonDetail); err != nil {
		return respPokeonInfo{}, err
	}

	return pokemonDetail, nil
}
