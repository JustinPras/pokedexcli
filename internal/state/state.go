package state

import (
	"github.com/JustinPras/pokedexcli/internal/pokeapi"
	"github.com/JustinPras/pokedexcli/internal/database"
)

type State struct {
	Db	*database.Queries
	Cfg	*Config
}


type Config struct {
	PokeapiClient pokeapi.Client
	PreviousURL *string
	NextURL *string
	Pokedex map[string]pokeapi.Pokemon
}