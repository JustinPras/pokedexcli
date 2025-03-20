package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)
	

func (c *Client) GetPokemon(pokemon string) (Pokemon, string, error) {
	url := baseURL + "/pokemon/" + pokemon

	if data, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		jsonStr := string(data)
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return Pokemon{}, "", err
		}
		return pokemonResp, jsonStr, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, "", err
	}

	resp, err := c.httpClient.Do(req)	
	if err != nil {
		return Pokemon{}, "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, "", err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, "", err
	}

	c.cache.Add(url, data)
	jsonStr := string(data)
	
	return pokemonResp, jsonStr, nil
}