package main

import (
	"fmt"
)

// implement this function
func commandExplore(config *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-areas>")
	}

	location := args[0]
	fmt.Printf("Exploring %s...\n", location)

	pokemonResp, err := config.pokeapiClient.ListPokemon(location)
	if err != nil {
		fmt.Println("Found no Pokemon")
		return err
	}


	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}