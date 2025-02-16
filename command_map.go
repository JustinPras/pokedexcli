package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationsResp.Next
	config.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

