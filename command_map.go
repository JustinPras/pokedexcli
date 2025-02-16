package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextURL, &config.pokeCache)
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

