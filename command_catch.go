package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(config *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}


	name := args[0]

	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		fmt.Println("Pokemon not found")
		return err
	}

	pokemonName := pokemon.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	userChance := r.Intn(1001)
	pokemonCatchDifficulty := pokemon.BaseExperience/10

	if userChance > pokemonCatchDifficulty {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	
	return nil
}