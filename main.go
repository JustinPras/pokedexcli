package main

import (
	"time"
	"log"
	"database/sql"
    
	_ "github.com/mattn/go-sqlite3"
    "github.com/JustinPras/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	pokedex := make(map[string]pokeapi.Pokemon)
	config := &Config{
		pokeapiClient: pokeClient,
		pokedex: pokedex,
	}
	startRepl(config)
}