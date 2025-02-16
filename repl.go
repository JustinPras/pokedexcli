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
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
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
	callback func(*Config) error
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

// type Config struct {
// 	previous int
// 	next int
// }

// func getConfig() *Config {
// 	return &Config{
// 		previous: 0,
// 		next: 21,
// 	}
// }

// func (config *Config) Next() {
// 	config.previous = config.next - 20
// 	config.next += 20
// }

// func (config *Config) Back() {
// 	config.next = config.previous + 20
// 	config.previous -= 20
// }

// type Location struct {
//     Name string `json:"name"`
// }
