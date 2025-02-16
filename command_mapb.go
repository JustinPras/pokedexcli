package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

func commandMapb(config *Config) error {
	if config.previous == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	for i := config.previous; i < config.previous + 20; i++ {
		resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", i))
		if err != nil {
			return fmt.Errorf("error making GET request - %s", err)
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var location Location
		err = json.Unmarshal(data, &location)
		if err != nil {
			return fmt.Errorf("error unmarshalling JSON - %s", err)
		}
		fmt.Println(location.Name)
	}
	config.Back()
	return nil
}