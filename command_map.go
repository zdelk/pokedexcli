package main

import (
	"errors"
	"fmt"
)

var offset int

func commandMapf(cfg *config, l string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, cfg.cache)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, l string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL, cfg.cache)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
