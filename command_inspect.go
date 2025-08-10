package main

import (
	"fmt"
)

func commandInspect(c *Config, a string) error {

	if len(a) == 0 {
		fmt.Println("Please specifiy a Pokemon")
		return nil
	}
	p, ok := c.pokeapiPokedex[a]
	if ok {
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Height: %v\n", p.Height)
		fmt.Printf("Weight: %v\n", p.Weight)
		fmt.Println("Stats:")
		for _, s := range p.Stats {
			fmt.Printf("  -%s: %v\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range p.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}

	} else {
		fmt.Printf("You haven't caught %s yet! Why not go try that first?\n", a)
	}
	return nil
}
