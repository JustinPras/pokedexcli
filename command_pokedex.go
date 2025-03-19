package main

import (
	"fmt"
)

func commandPokedex(s *state, args []string) error{
	if len(s.cfg.pokedex) == 0 {
		return fmt.Errorf("You have caught no Pokemon")
	}

	fmt.Print("Your Pokedex:\n")
	for _, pokemon := range(s.cfg.pokedex) {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}