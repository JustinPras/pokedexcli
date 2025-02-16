package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)



func commandMap(config *Config) error {
	fmt.Println("config.previous", config.previous)
	fmt.Println("config.next", config.next)
	for i := config.next - 20; i < config.next; i++ {
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
	config.Next()
	return nil
}

