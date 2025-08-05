package main

import (
	"fmt"
)

func commandExplore(c *Config, a string) error {
	if len(a) == 0 {
		fmt.Println("Please specifiy a location")
		return nil
	}

	areaResp, err := c.pokeapiClient.ExploreLocation(a)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", a)
	fmt.Println("Found Pokemon:")
	for _, r := range areaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", r.Pokemon.Name)
	}
	// fmt.Println(areaResp)

	return nil
}
