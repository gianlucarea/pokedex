package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func cmdExit(_ *config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cmdHelp(_ *config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func cmdMap(cfg *config, _ []string) error {
	response, err := cfg.apiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = response.Next
	cfg.previous = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func cmdMapb(cfg *config, _ []string) error {
	if cfg.previous == nil {
		return errors.New("This is first page")
	}
	response, err := cfg.apiClient.ListLocations(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = response.Next
	cfg.previous = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func cmdExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location area name")
	}
	name := args[0]
	fmt.Printf("Exploring %s...\n", name)

	resp, err := cfg.apiClient.GetLocationArea(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, enc := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}

func cmdCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.apiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	threshold := rand.Intn(pokemon.BaseExperience)
	if threshold > 40 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.pokedex[name] = pokemon
	return nil
}

func cmdInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
