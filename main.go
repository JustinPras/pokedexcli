package main

import (
	"time"
	"github.com/JustinPras/pokedexcli/internal/pokeapi"
	"github.com/JustinPras/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	config := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}