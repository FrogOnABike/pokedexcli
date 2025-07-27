package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/frogonabike/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Displays previous page of locations",
			callback:    commandMapb,
		},
	}
}

func startRepl(c *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		userInput := reader.Text()
		userWords := cleanInput(userInput)
		if len(userWords) == 0 {
			fmt.Println("Please enter a command")
			continue
		}

		command, ok := getCommand()[userWords[0]]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}

		err := command.callback(c)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

	}
}

// Format user input to all lower case and split by words, returning a slice

func cleanInput(text string) []string {
	lowerWords := strings.ToLower((text))
	words := strings.Fields(lowerWords)
	return words
}
