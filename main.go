package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}
		text := scanner.Text()
		input := cleanInput(text)
		fmt.Printf("Your command was: %v\n", input[0])
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Fields(lower)
	return split
}