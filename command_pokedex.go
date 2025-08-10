package main

import (
	"fmt"
)

func commandPokedex(c *Config, a string) error {

	if len(c.pokeapiPokedex) == 0 {
		fmt.Println("Your Pokedex appears to be empty :( Go try catching something!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, p := range c.pokeapiPokedex {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
