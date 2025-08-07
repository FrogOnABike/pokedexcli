package main

import (
	"fmt"
)

func commandCatch(c *Config, a string) error {
	if len(a) == 0 {
		fmt.Println("Please specifiy a Pokemon")
		return nil
	}

	pokemonResp, err := c.pokeapiClient.CatchPokemon(a)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s\n", a)
	fmt.Printf("Base experience: %v\n", pokemonResp.BaseExperience)
	// fmt.Println(pokemonResp)
	return nil
}
