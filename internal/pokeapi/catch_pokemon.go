package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (PokemonInfo, error) {
	pageURL := baseURL + "/pokemon/" + pokemonName

	cacheData, exists := c.clientCache.Get(pageURL)
	if exists {
		fmt.Println("***Retrieving from cache***")
		pokemonDetail := PokemonInfo{}

		if err := json.Unmarshal(cacheData, &pokemonDetail); err != nil {
			return PokemonInfo{}, err
		}
		return pokemonDetail, nil
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokemonDetail := PokemonInfo{}
	c.clientCache.Add(pageURL, body)

	if err := json.Unmarshal(body, &pokemonDetail); err != nil {
		return PokemonInfo{}, err
	}

	return pokemonDetail, nil
}
