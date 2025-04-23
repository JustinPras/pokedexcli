package commands

import (
	"github.com/JustinPras/pokedexcli/internal/state"
)

var gameVer = "platinum"

type cliCommand struct {
	name string
	description string
	Callback func(*state.State, []string) error
} 

type OrderedCommands struct {
	Order 		[]string
	CommandMap 	map[string]cliCommand
}

func GetCommands() OrderedCommands {
	commands :=  map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next page of locations in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations in the Pokemon world",
			Callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore the Pokemon world",
			Callback:    commandExplore,
		},
		"catch": {
			name:		"catch <pokemon_name>",
			description:"Throw a Pokeball at a pokemon",
			Callback: commandCatch,
		},
		"inspect": {
			name: 		"inspect <pokemon_name>",
			description:"Inspect a Pokemon from your Pokedex",
			Callback:	commandInspect,
		},
		"pokedex": {
			name:		"pokedex",
			description:"Displays all Pokemon in your Pokedex",
			Callback:	commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}

	order := []string{"help", "map", "mapb", "explore", "catch", "inspect", "pokedex", "exit"}

	return OrderedCommands{
        Order:      order,
        CommandMap: commands,
    }
}