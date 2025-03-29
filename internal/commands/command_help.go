package commands

import (
	"fmt"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func commandHelp(s *state.State, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}