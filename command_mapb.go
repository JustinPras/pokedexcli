package main

import (
	"fmt"
	"errors"
)

func commandMapb(state *state, args []string) error {
	if state.cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := state.cfg.pokeapiClient.ListLocations(state.cfg.previousURL)
	if err != nil {
		return err
	}

	state.cfg.nextURL = locationsResp.Next
	state.cfg.previousURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}