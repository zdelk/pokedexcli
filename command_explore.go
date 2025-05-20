package main

import (
	"fmt"
)

func commandExplore(cfg *config, location string) error {
	fmt.Printf("Exploring %s...\n", location)
	fullInfo, err := cfg.pokeapiClient.ListPokemon(location, cfg.cache)
	if err != nil {
		return err
	}
	pokeList := fullInfo.PokemonEncounters
	fmt.Println("Found Pokemon:")
	for _, mon := range pokeList {
		pokemon := mon.Pokemon.Name
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}
