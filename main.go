package main

import (
	"time"

	"github.com/zmdelk/pokedexcli/internal/pokeapi"
	"github.com/zmdelk/pokedexcli/internal/pokecache"
)

func main() {
	pokeCache := pokecache.NewCache(5 * time.Second)

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         pokeCache,
	}
	startRepl(cfg)
}
