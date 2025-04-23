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

	cmds := GetCommands()

	for _, commandName := range cmds.Order {
		cmd := cmds.CommandMap[commandName]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}