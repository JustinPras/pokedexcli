package main

import (
	"fmt"
	"math/rand"
	"time"
	"context"

	"github.com/JustinPras/pokedexcli/internal/database"
	"github.com/JustinPras/pokedexcli/internal/pokeapi"
)

func commandCatch(s *state, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	name := args[0]
	pokemon, err := s.cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	pokemonName := pokemon.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	userChance := r.Intn(pokemon.BaseExperience)

	if userChance <= 40 {
		fmt.Printf("%s was caught!\n", pokemonName)
		s.cfg.pokedex[pokemonName] = pokemon
		err = createPokedexEntry(s, pokemon)
		if err != nil {
			return fmt.Errorf("Error storing pokemon in database: %w", err)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	
	return nil
}

func createPokedexEntry(s *state, pokemon pokeapi.Pokemon) error {
	createPokedexEntryParams := database.CreatePokedexEntryParams {
		PokemonName:	pokemon.Name,
		Experience:		int64(pokemon.BaseExperience),
		CapturedAt:		time.Now(),
		Height:			int64(pokemon.Height),
		Weight:			int64(pokemon.Weight),
		PokemonID:		int64(pokemon.ID),
	}

	_, err := s.db.CreatePokedexEntry(context.Background(), createPokedexEntryParams)
	if err != nil {
		return fmt.Errorf("Error creating Pokedex Entry: %w", err)
	}	

	return nil
}