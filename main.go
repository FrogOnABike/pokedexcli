package main

import (
	"time"

	"github.com/frogonabike/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	c := Config{
		pokeapiClient: *pokeClient,
	}
	startRepl(&c)
}
