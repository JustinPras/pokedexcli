package main

import (
	"fmt"
	"os"
)

func commandExit(state *state, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}