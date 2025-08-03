package main

import (
	"fmt"
	"os"
)

func commandExplore(c *Config, a string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
