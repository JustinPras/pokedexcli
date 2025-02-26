package main

import (
	"fmt"
)

func commandMap(config *Config, args []string) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextURL)
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

