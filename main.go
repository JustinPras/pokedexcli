package main

import (
	"time"
	"log"
	"database/sql"
	"encoding/json"
	"context"
    
	_ "github.com/mattn/go-sqlite3" //encryption sqlite driver
	// _ "modernc.org/sqlite" //pure go implementation of sqlite
    "github.com/JustinPras/pokedexcli/internal/pokeapi"
	"github.com/JustinPras/pokedexcli/internal/database"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	pokedex := make(map[string]pokeapi.Pokemon)

	config := &state.Config{
		PokeapiClient: pokeClient,
		Pokedex: pokedex,
	}

	db, err := sql.Open("sqlite3", "pokedex.db")
	if err != nil {
		log.Fatal("Error connecting to the database: %w", err)
	}

	dbQueries := database.New(db)

	programState := state.State {
		Db:		dbQueries,
		Cfg:	config,
	}

	fillPokedex(&programState)
	startRepl(&programState)
}


func fillPokedex(s *state.State) error {
	pokedex, err := s.Db.GetPokedex(context.Background())
	if err != nil {
		return err
	}

	pokemon := pokeapi.Pokemon{}

	for _, pokedexEntry := range(pokedex) {
		err = json.Unmarshal([]byte(pokedexEntry.JsonData), &pokemon)
		if err != nil {
			return err
		}
		s.Cfg.Pokedex[pokedexEntry.PokemonName] = pokemon 
	}
	return nil
}