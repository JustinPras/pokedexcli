package main

import (
	"fmt"
)

func commandInspect(config *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	name := args[0]

	pokemon, ok := config.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %b\n", pokemon.Height)
	fmt.Printf("Weight: %b\n", pokemon.Weight)
	fmt.Print("Stats:\n")
	for _, stat := range(pokemon.Stats) {
		fmt.Printf("\t-%s: %b\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, typef := range(pokemon.Types) {
		fmt.Printf("\t- %s\n", typef.Type.Name)
	}

	return nil
}