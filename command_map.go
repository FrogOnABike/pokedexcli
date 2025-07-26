package main

import "fmt"

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
