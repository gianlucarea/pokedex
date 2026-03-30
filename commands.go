package main

import (
	"errors"
	"fmt"
	"os"
)

func cmdExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cmdHelp(_ *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func cmdMap(cfg *config) error {
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

func cmdMapb(cfg *config) error {
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
