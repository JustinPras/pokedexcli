package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"github.com/JustinPras/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.Client
	previousURL *string
	nextURL *string
	pokedex map[string]pokeapi.Pokemon
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}
		
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		args := text[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unkown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func(*Config, []string) error
} 

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next page of locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the Pokemon world",
			callback:    commandExplore,
		},
		"catch": {
			name:		"catch",
			description: "Throw a Pokeball at a pokemon",
			callback: commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}