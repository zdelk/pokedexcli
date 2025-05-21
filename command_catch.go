package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemon := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	pokeInfo, err := cfg.pokeapiClient.CatchRate(pokemon)
	if err != nil {
		return err
	}
	baseEXP := pokeInfo.BaseExperience
	catchNum := rand.Intn(256)
	if catchNum < baseEXP {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon)
	cfg.pokedex[pokemon] = pokeInfo
	fmt.Println("You may now inspect it with the inspect command.")

	return nil
}
