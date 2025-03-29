package commands

import (
	"fmt"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func commandInspect(s *state.State, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	name := args[0]

	pokemon, ok := s.Cfg.Pokedex[name]
	if !ok {
		return fmt.Errorf("You have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %b\n", pokemon.Height)
	fmt.Printf("Weight: %b\n", pokemon.Weight)
	fmt.Print("Stats:\n")
	for _, stat := range(pokemon.Stats) {
		fmt.Printf("  -%s: %b\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, typef := range(pokemon.Types) {
		fmt.Printf("  - %s\n", typef.Type.Name)
	}

	return nil
}