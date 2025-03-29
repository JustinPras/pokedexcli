package commands

import (
	"fmt"
	"math/rand"
	"time"
	"context"

	"github.com/JustinPras/pokedexcli/internal/database"
	"github.com/JustinPras/pokedexcli/internal/pokeapi"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func commandCatch(s *state.State, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	name := args[0]
	pokemon, jsonStr, err := s.Cfg.PokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	pokemonName := pokemon.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	userChance := r.Intn(pokemon.BaseExperience)

	if userChance <= 40 {
		fmt.Printf("%s was caught!\n", pokemonName)
		s.Cfg.Pokedex[pokemonName] = pokemon
		err = createPokedexEntry(s, pokemon, jsonStr)
		if err != nil {
			return fmt.Errorf("Error storing pokemon in database: %w", err)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	
	return nil
}

func createPokedexEntry(s *state.State, pokemon pokeapi.Pokemon, jsonStr string) error {
	createPokedexEntryParams := database.CreatePokedexEntryParams {
		PokemonName:	pokemon.Name,
		PokemonID:		int64(pokemon.ID),
		JsonData:		jsonStr,
	}

	_, err := s.Db.CreatePokedexEntry(context.Background(), createPokedexEntryParams)
	if err != nil {
		return fmt.Errorf("Error creating Pokedex Entry: %w", err)
	}	

	return nil
}