package commands

import (
	"fmt"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func commandPokedex(s *state.State, args []string) error{
	if len(s.Cfg.Pokedex) == 0 {
		return fmt.Errorf("You have caught no Pokemon")
	}

	fmt.Print("Your Pokedex:\n")
	for _, pokemon := range(s.Cfg.Pokedex) {
		fmt.Printf(" - %s\n", pokemon.Name)
	}


	return nil
}