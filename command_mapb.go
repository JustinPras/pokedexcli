package main

import (
	"fmt"
	"errors"
)

func commandMapb(s *state, args []string) error {
	if s.cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := s.cfg.pokeapiClient.ListLocations(s.cfg.previousURL)
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