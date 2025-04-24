package commands

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/JustinPras/pokedexcli/internal/state"
	"github.com/JustinPras/pokedexcli/internal/pokeapi"
)

var encounterMethods = []string{"walk", "dark-grass", "grass-spots", "cave-spots"} // depends on type of gameplay

// implement this function
func commandExplore(s *state.State, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-areas>")
	}

	name := args[0]
	

	location, err := s.Cfg.PokeapiClient.GetLocation(name)
	if err != nil {
		fmt.Println("Found no Pokemon")
		return err
	}

	fmt.Printf("\nExploring %s...\n\n", location.Name)

	pokemonName, ok := encounter(location)
	if ok {
		fmt.Printf("You encountered a %s!\n\n", pokemonName)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	fmt.Println()
	return nil
}

func encounter(location pokeapi.Location) (string, bool) {
	if len(location.PokemonEncounters) != 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, encounter := range location.PokemonEncounters {
			for _, versionDetails := range encounter.VersionDetails {
				if versionDetails.Version.Name == gameVer {
					for _, encounterDetails := range versionDetails.EncounterDetails {
						if contains(encounterMethods, encounterDetails.Method.Name) {
							userChance := r.Intn(100)
							if userChance < encounterDetails.Chance {
								return encounter.Pokemon.Name, true
							}
						}
					}
					
				}
			}
		}
	}
	return "", false
}

func contains(slice []string, item string) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}