package main

import (
	"time"
	"log"
	"database/sql"
	"fmt"
	"encoding/json"
	"context"
    
	_ "github.com/mattn/go-sqlite3"
    "github.com/JustinPras/pokedexcli/internal/pokeapi"
	"github.com/JustinPras/pokedexcli/internal/database"
)

type state struct {
	db	*database.Queries
	cfg	*Config
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	pokedex := make(map[string]pokeapi.Pokemon)

	config := &Config{
		pokeapiClient: pokeClient,
		pokedex: pokedex,
	}

	fmt.Println("I made it to the beginning of the DB code!")

	db, err := sql.Open("sqlite3", "pokedex.db")
	if err != nil {
		log.Fatal("Error connecting to the database: %w", err)
	}

	dbQueries := database.New(db)

	programState := state {
		db:		dbQueries,
		cfg:	config,
	}

	fillPokedex(&programState)

	fmt.Println("I made it to the end of main!")
	startRepl(&programState)
}


func fillPokedex(s *state) error {
	pokedex, err := s.db.GetPokedex(context.Background())
	if err != nil {
		return err
	}

	pokemon := pokeapi.Pokemon{}

	for _, pokedexEntry := range(pokedex) {
		err = json.Unmarshal([]byte(pokedexEntry.JsonData), &pokemon)
		if err != nil {
			return err
		}
		s.cfg.pokedex[pokedexEntry.PokemonName] = pokemon 
	}
	return nil
}