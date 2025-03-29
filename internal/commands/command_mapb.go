package commands

import (
	"fmt"
	"errors"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func commandMapb(s *state.State, args []string) error {
	if s.Cfg.PreviousURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := s.Cfg.PokeapiClient.ListLocations(s.Cfg.PreviousURL)
	if err != nil {
		return err
	}

	s.Cfg.NextURL = locationsResp.Next
	s.Cfg.PreviousURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}