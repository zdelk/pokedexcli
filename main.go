package main

import (
	"time"

	"github.com/zmdelk/pokedexcli/internal/pokeapi"
	"github.com/zmdelk/pokedexcli/internal/pokecache"
)

func main() {
	pokeCache := pokecache.NewCache(5 * time.Second)

	pokeClient := pokeapi.NewClient(5 * time.Second)
	usrPokedex := make(map[string]pokeapi.PokeStats)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         pokeCache,
		pokedex:       usrPokedex,
	}
	startRepl(cfg)
}
