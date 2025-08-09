package main

import (
	"time"

	"github.com/frogonabike/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 30*time.Second)
	usrPokedex := make(pokeapi.Pokedex)
	c := Config{
		pokeapiClient:  *pokeClient,
		pokeapiPokedex: usrPokedex,
	}
	startRepl(&c)
}
