package main

import (
	"fmt"
	"math/rand"
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

	// Roll a virtual D100 to see what the players catch factor is :)
	catchChance := rand.Intn(100)

	// Set a default catch chance of 35%
	target := 35

	fmt.Printf("Throwing a Pokeball at %s...\n", a)
	fmt.Printf("**DEBUG** Catch Chance:%v\n", catchChance)
	fmt.Printf("Base experience: %v\n", pokemonResp.BaseExperience)

	// Set catch chance based on target Pokemon's base experience
	switch {
	case pokemonResp.BaseExperience >= 500:
		target = 95
	case pokemonResp.BaseExperience >= 400:
		target = 85
	case pokemonResp.BaseExperience >= 300:
		target = 75
	case pokemonResp.BaseExperience >= 200:
		target = 65
	case pokemonResp.BaseExperience >= 100:
		target = 55
	}

	// Compare players catch factor to the target, did they get it?
	if catchChance >= target {
		fmt.Printf("%s was caught!\n", a)
		c.pokeapiPokedex[a] = pokemonResp
		fmt.Printf("You have caught %v Pokemon!\n", len(c.pokeapiPokedex))
	} else {
		fmt.Printf("%s escaped :(\n", a)
		fmt.Printf("You have caught %v Pokemon!\n", len(c.pokeapiPokedex))
	}
	// fmt.Println(pokemonResp)
	return nil
}
