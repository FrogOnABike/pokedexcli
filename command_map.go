package main

import (
	"errors"
	"fmt"
)

func commandMap(c *Config) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.nextURL)
	if err != nil {
		return err
	}
	c.nextURL = &locationsResp.Next
	c.prevURL = &locationsResp.Previous

	for _, r := range locationsResp.Results {
		fmt.Println(r.Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	if *c.prevURL == "" {
		return errors.New("you're on the first page")
	}

	locationsResp, err := c.pokeapiClient.ListLocations(c.prevURL)
	if err != nil {
		return err
	}
	c.nextURL = &locationsResp.Next
	c.prevURL = &locationsResp.Previous

	for _, r := range locationsResp.Results {
		fmt.Println(r.Name)
	}

	return nil
}
