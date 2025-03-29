package commands

import (
	"fmt"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func commandMap(s *state.State, args []string) error {
	locationsResp, err := s.Cfg.PokeapiClient.ListLocations(s.Cfg.NextURL)
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

