package commands

import (
	"fmt"
	"os"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func commandExit(s *state.State, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}