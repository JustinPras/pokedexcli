package main

import (
	"fmt"
	"context"

	"github.com/JustinPras/pokedexcli/internal/pokeapi"	
)

func commandInspect(s *state, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	name := args[0]

	pokemon, err := getPokemonByName(s, name)
	if err != nil {
		return fmt.Errorf("you have not caught that pokemon")
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

func getPokemonByName(s *state, pokemonName string) (pokeapi.Pokemon, error) {
	pokedexEntry, err := s.db.GetPokemonByName(context.Background(), pokemonName)
	//handle if the pokemone isn't in the pokedex (hasn't been caught)
	if err != nil {
		return pokeapi.Pokemon{}, fmt.Errorf("Error retrieving pokemon: %w", err)
	}

	pokemon, err := s.cfg.pokeapiClient.GetPokemon(pokedexEntry.PokemonName)
	if err != nil {
		return pokeapi.Pokemon{}, err
	}

	return pokemon, nil
}