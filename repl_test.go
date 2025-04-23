package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "charmander bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual: %q, expected: %q", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual: %q, expected: %q", actual, c.expected)
			}
			
		}
	}
}

func TestHelpCommand(t *testing.T) {
	cmd := exec.Command("go", "run", ".")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out


	cmd.Stdin = bytes.NewReader([]byte("help\n"))
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Error running command: %v", err)
	}


	expectedOutput := 	`Pokedex > 
Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Display the next page of locations in the Pokemon world
mapb: Display the previous page of locations in the Pokemon world
explore <location_name>: Explore the Pokemon world
catch <pokemon_name>: Throw a Pokeball at a pokemon
inspect <pokemon_name>: Inspect a Pokemon from your Pokedex
pokedex: Displays all Pokemon in your Pokedex
exit: Exit the Pokedex

Pokedex > `

	actualOutput := out.String()

	if actualOutput != expectedOutput {
		t.Errorf("Help command output is incorrect.\nGot:\n%s\nExpected:\n%s", actualOutput, expectedOutput)
	}
}