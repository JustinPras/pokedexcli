package main

import (
	"fmt"
	"strings"
	"io"
	"github.com/JustinPras/pokedexcli/internal/commands"
	"github.com/JustinPras/pokedexcli/internal/state"

	"github.com/peterh/liner"
)

func startRepl(state *state.State) {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	for {

		input, err := line.Prompt("Pokedex > ")
        if err == liner.ErrPromptAborted || err == io.EOF {
            break
        }
		
		line.AppendHistory(input)

		text := cleanInput(input)
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