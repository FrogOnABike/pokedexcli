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
	callback    func(c *Config, a string) error
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
			name:        "mapb",
			description: "Displays previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays Pokemon that can be found at the specified location",
			callback:    commandExplore,
		},
	}
}

func startRepl(c *Config) {
	arg := ""
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		userInput := reader.Text()
		userWords := cleanInput(userInput)

		//check if a command passed by confirming length of returned words map - if 0 then no command supplied
		if len(userWords) == 0 {
			fmt.Println("Please enter a command")
			continue
		}

		//check command is valie by passing to getCommand and capturing err if raised
		command, ok := getCommand()[userWords[0]]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}

		//check if extra arguments passed - for example if supplying a location name or pokemon - only using first string after the command (userWords[1] item at present)
		if len(userWords) >= 2 {
			arg = userWords[1]
		}

		err := command.callback(c, arg)
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
