package main

import (
	"time"
	"github.com/JustinPras/pokedexcli/internal/pokeapi"
	"github.com/JustinPras/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	config := &Config{
		pokeapiClient: pokeClient,
		pokeCache: *pokeCache,
	}
	startRepl(config)
}