package main

import (
	"fmt"
)

func commandPokedex(config *Config, args []string) error{
	if len(config.pokedex) == 0 {
		return fmt.Errorf("You have caught no Pokemon")
	}

	fmt.Print("Your Pokedex:\n")
	for _, pokemon := range(config.pokedex) {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}