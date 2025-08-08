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
	fmt.Printf("Throwing a Pokeball at %s\n", a)
	fmt.Printf("Base experience: %v\n", pokemonResp.BaseExperience)
	catchChance := rand.Intn(100)
	target := 35

	switch {
	case pokemonResp.BaseExperience >= 500:
		target = 95
		// if catchChance >= 95 {
		// caught = true
		// break
		// }
	case pokemonResp.BaseExperience >= 400:
		target = 85
		// if catchChance >= 85 {
		// 	caught = true
		// 	break
		// }
	case pokemonResp.BaseExperience >= 300:
		target = 75
		// if catchChance >= 75 {
		// 	caught = true
		// 	break
		// }
	case pokemonResp.BaseExperience >= 200:
		target = 65
		// if catchChance >= 65 {
		// 	caught = true
		// 	break
		// }
	case pokemonResp.BaseExperience >= 100:
		target = 55
		// if catchChance >= 55 {
		// 	caught = true
		// 	break
		// }
	}

	if catchChance >= target {
		fmt.Printf("%s was caught!\n", a)
	} else {
		fmt.Printf("%s escaped :(\n", a)
	}

	// fmt.Println(pokemonResp)
	return nil
}
