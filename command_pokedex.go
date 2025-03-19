package main

import (
	"fmt"
)

func commandPokedex(state *state, args []string) error{
	if len(state.cfg.pokedex) == 0 {
		return fmt.Errorf("You have caught no Pokemon")
	}

	fmt.Print("Your Pokedex:\n")
	for _, pokemon := range(state.cfg.pokedex) {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}