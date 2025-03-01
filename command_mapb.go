package main

import (
	"fmt"
	"errors"
)

func commandMapb(config *Config, args []string) error {
	if config.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := config.pokeapiClient.ListLocations(config.previousURL)
	if err != nil {
		return err
	}

	config.nextURL = locationsResp.Next
	config.previousURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}