package main

import (
	"fmt"
)

// implement this function
func commandExplore(state *state, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-areas>")
	}

	name := args[0]
	

	location, err := state.cfg.pokeapiClient.GetLocation(name)
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