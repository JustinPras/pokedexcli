package commands

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/JustinPras/pokedexcli/internal/state"
)

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

	fmt.Printf("Exploring %s...\n", location.Name)

	encounterMethods := []string{"walk", "dark-grass", "grass-spots", "cave-spots"}

	
	if len(location.PokemonEncounters) != 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		outerLoop:
		for _, encounter := range location.PokemonEncounters {
			for _, versionDetails := range encounter.VersionDetails {
				if versionDetails.Version.Name == gameVer {
					for _, encounterDetails := range versionDetails.EncounterDetails {
						if contains(encounterMethods, encounterDetails.Method.Name) {
							userChance := r.Intn(100)
							if userChance < encounterDetails.Chance {
								fmt.Printf("You encountered a %s\n", encounter.Pokemon.Name)
								break outerLoop;
							}
						}
					}
					
				}
			}
		}
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func contains(slice []string, item string) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}