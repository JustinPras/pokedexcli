package main

import (
	"time"
	"log"
	"database/sql"
    
	_ "github.com/mattn/go-sqlite3"
    "github.com/JustinPras/pokedexcli/internal/pokeapi"
)

type state struct {
	db	*database.Queries
	cfg	*config.Config
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	pokedex := make(map[string]pokeapi.Pokemon)
	
	config := &Config{
		pokeapiClient: pokeClient,
		pokedex: pokedex,
	}

	db, err := sql.Open("sqlite3", "pokedex.db")
	if err != nil {
		log.Fatal("Error connecting to the database: %w", err)
	}

	dbQueries := database.New(db)

	programState := state {
		db:		dbQueries,
		cfg:	config,
	}

	startRepl(programState)
}