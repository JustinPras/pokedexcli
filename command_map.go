package main

import (
	"fmt"
)

func commandMap(s *state, args []string) error {
	locationsResp, err := s.cfg.pokeapiClient.ListLocations(s.cfg.nextURL)
	if err != nil {
		return err
	}

	s.cfg.nextURL = locationsResp.Next
	s.cfg.previousURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

