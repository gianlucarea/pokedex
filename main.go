package main

import (
	"time"

	"github.com/gianlucarea/pokedex/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		apiClient: pokeClient,
	}
	Start(cfg)
}
