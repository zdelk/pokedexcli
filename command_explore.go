package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	location := args[0]

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
