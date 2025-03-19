package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(state *state, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	name := args[0]
	pokemon, err := state.cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	pokemonName := pokemon.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	userChance := r.Intn(pokemon.BaseExperience)

	if userChance <= 40 {
		fmt.Printf("%s was caught!\n", pokemonName)
		state.cfg.pokedex[pokemonName] = pokemon
		err = createPokedexEntry(pokemon)
		if err != nil {
			return fmt.Errorf("Error storing pokemon in database: %w", err)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	
	return nil
}

func createPokedexEntry(pokemon pokeapi.Pokemon) error {

}