package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var offset int

func commandMap() error {
	apiLink := "https://pokeapi.co/api/v2/location-area/?limit=20"
	if offset != 0 {
		apiLink = apiLink + fmt.Sprintf("&offset=%d", offset)
	}
	res, err := http.Get(apiLink)
	if err != nil {
		return fmt.Errorf("error in request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error in reading: %v", err)
	}
	var locations PokemonAreas

	if err := json.Unmarshal(body, &locations); err != nil {
		return fmt.Errorf("error in unmarshal: %v", err)
	}
	areas := locations.Results

	for _, area := range areas {
		fmt.Println(area.Name)
	}
	offset += 20
	return nil
}

// func commandMapb() {
// 	offset -= 20
// 	apiLink := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offest=%d",offset)
// }

type PokemonAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
