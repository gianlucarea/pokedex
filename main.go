package main

import (
	"time"

	"github.com/gianlucarea/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		apiClient: pokeClient,
	}
	Start(cfg)
}
