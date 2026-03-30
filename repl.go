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
	pokedex   map[string]pokeapi.RespPokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
	"explore": {
		name:        "explore",
		description: "Explore Locations",
		callback:    cmdExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch a Pokemon",
		callback:    cmdCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect a caught Pokemon",
		callback:    cmdInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "List all caught Pokemon",
		callback:    cmdPokedex,
	},
}

func Start(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			words := CleanInput(text)

			cmd, ok := supportedCMD[words[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			cmd.callback(cfg, words[1:])
		}
	}
}

func CleanInput(text string) []string {
	return strings.Fields(text)
}
