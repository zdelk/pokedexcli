package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must enter a pokemon")
	}
	pokemon := args[0]
	info, ok := cfg.pokedex[pokemon]
	if !ok {
		return fmt.Errorf("you haven't caught a %s yet", pokemon)
	}

	fmt.Printf("Name: %s\n", pokemon)
	fmt.Printf("Height: %d\n", info.Height)
	fmt.Printf("Weight: %d\n", info.Weight)
	fmt.Println("Stats:")
	for _, stat := range info.Stats {
		fmt.Printf(" -%s:  %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ptype := range info.Types {
		fmt.Printf(" -%s\n", ptype.Type.Name)
	}
	return nil
}
