package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"github.com/JustinPras/pokedexcli/internal/commands"
	"github.com/JustinPras/pokedexcli/internal/state"
)

func startRepl(state *state.State) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}
		
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		args := text[1:]

		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Callback(state, args)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unkown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}