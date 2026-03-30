package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gianlucarea/pokedex/internal/pokeapi"
)

type config struct {
	apiClient pokeapi.Client
	next      *string
	previous  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var supportedCMD map[string]cliCommand = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    cmdExit,
	},
	"help": {
		name:        "help",
		description: "Help Function of the Pokedex",
		callback:    cmdHelp,
	},
	"map": {
		name:        "map",
		description: "Next Map Locations",
		callback:    cmdMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Previous Map Locations",
		callback:    cmdMapb,
	},
}

func Start(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())
		cmd, ok := supportedCMD[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback(cfg)
	}
}

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
