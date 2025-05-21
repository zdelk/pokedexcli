package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) < 1 {
		return errors.New("you haven't caught any pokemon")
	}
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.pokedex {
		fmt.Println(" - ", key)
	}
	return nil
}
