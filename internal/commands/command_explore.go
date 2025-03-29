package commands

import (
	"fmt"
	"github.com/JustinPras/pokedexcli/internal/state"
)

// implement this function
func commandExplore(s *state.State, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-areas>")
	}

	name := args[0]
	

	location, err := s.Cfg.PokeapiClient.GetLocation(name)
	if err != nil {
		fmt.Println("Found no Pokemon")
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)

	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}