package main

import "github.com/frogonabike/pokedexcli/internal/pokeapi"

func main() {
	pokeClient := pokeapi.NewClient()
	c := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(c)
}
