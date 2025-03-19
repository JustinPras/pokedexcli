package main

import (
	"fmt"
)

func commandMap(state *state, args []string) error {
	locationsResp, err := state.cfg.pokeapiClient.ListLocations(state.cfg.nextURL)
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

